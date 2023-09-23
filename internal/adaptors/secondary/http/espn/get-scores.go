package espn

import (
	"cricwatch/internal/core/domain"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
)

var (
	getScorePath                = cacheID + "/summary"
	ErrUnknownResponseStructure = errors.New("unknown response structure")
)

type Commentary struct {
	SortOrder float32 `json:"sortOrder"`
	HomeScore string  `json:"homeScore"`
	AwayScore string  `json:"awayScore"`
}

type Competition struct {
	Description string `json:"description"`
	// Commentaries []Commentary `json:"commentaries"`
	Competitors []Competitor `json:"competitors"`
}

type Competitor struct {
	HomeAway string `json:"homeAway"`
	Team     struct {
		Name string `json:"name"`
	}
	LineScores []LineScore `json:"linescores"`
}

func (c *Competitor) toTeamScore() (domain.TeamScore, error) {
	if len(c.LineScores) == 2 {
		if c.LineScores[0].Runs == 0 && c.LineScores[1].Runs > 0 {
			c.LineScores = []LineScore{c.LineScores[1]}
		} else if c.LineScores[0].Runs > 0 && c.LineScores[1].Runs == 0 {
			c.LineScores = []LineScore{c.LineScores[0]}
		} else {
			c.LineScores = []LineScore{c.LineScores[0]}
		}
	}

	if len(c.LineScores) != 1 {
		return domain.TeamScore{}, errors.Join(ErrUnknownResponseStructure, fmt.Errorf("got %d line scores, want 1", len(c.LineScores)))
	}

	unparsedOvers := c.LineScores[0].Overs
	overs := int(math.Floor(unparsedOvers))
	balls := int(math.Round((unparsedOvers - math.Floor(unparsedOvers)) * 10))

	return domain.TeamScore{
		Team:    domain.NewTeam(c.Team.Name),
		Runs:    domain.NewRuns(c.LineScores[0].Runs),
		Wickets: domain.NewWickets(c.LineScores[0].Wickets),
		Overs:   domain.NewOvers(overs, balls),
	}, nil
}

func (c *Competitor) IsHome() bool {
	return c.HomeAway == "home"
}

func (c *Competitor) IsAway() bool {
	return c.HomeAway == "away"
}

type LineScore struct {
	Runs    int     `json:"runs"`
	Wickets int     `json:"wickets"`
	Overs   float64 `json:"overs"`
}

type GetScoreResponse struct {
	Header struct {
		Name         string        `json:"name"`
		Competitions []Competition `json:"competitions"`
	} `json:"header"`
}

func (resp *GetScoreResponse) toScore() (domain.Score, error) {
	var score domain.Score

	for _, competition := range resp.Header.Competitions {
		for _, competitor := range competition.Competitors {
			if competitor.IsHome() {
				s, err := competitor.toTeamScore()
				if err != nil {
					return domain.Score{}, err
				}
				score.Home = s
			} else if competitor.IsAway() {
				s, err := competitor.toTeamScore()
				if err != nil {
					return domain.Score{}, err
				}
				score.Away = s
			}
		}
	}

	return score, nil
}

func (c *Client) GetScore(match domain.Match) (domain.Score, error) {
	url, err := url.JoinPath(baseUrl, getScorePath)
	if err != nil {
		return domain.Score{}, err
	}

	req, err := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("event", fmt.Sprintf("%d", match.ID))
	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return domain.Score{}, err
	}
	if resp.StatusCode != 200 {
		return domain.Score{}, fmt.Errorf("error while fetching score, status code: %d", resp.StatusCode)
	}

	var getScoreResponse GetScoreResponse

	err = json.NewDecoder(resp.Body).Decode(&getScoreResponse)
	if err != nil {
		return domain.Score{}, err
	}

	return getScoreResponse.toScore()
}
