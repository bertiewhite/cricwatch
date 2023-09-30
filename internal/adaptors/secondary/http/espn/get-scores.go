package espn

import (
	"cricwatch/internal/core/domain"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
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
	Description string       `json:"description"`
	Competitors []Competitor `json:"competitors"`
}

func (c *Competition) GetBattingTeam() Competitor {
	for _, competitor := range c.Competitors {
		if competitor.isBatting() {
			return competitor
		}
	}
	return Competitor{}
}

func (c *Competition) GetBowlingTeam() Competitor {
	for _, competitor := range c.Competitors {
		if !competitor.isBatting() {
			return competitor
		}
	}
	return Competitor{}
}

type Competitor struct {
	HomeAway string `json:"homeAway"`
	Team     struct {
		Name string `json:"name"`
	}
	LineScores []LineScore `json:"linescores"`
}

func (c *Competitor) isBatting() bool {
	latestLinescore := c.LineScores[len(c.LineScores)-1]
	return latestLinescore.IsBatting
}

func (c *Competitor) toTeamScore() (domain.TeamScore, error) {
	lineScore := LineScore{
		Runs:      0,
		Wickets:   0,
		Overs:     0.0,
		IsBatting: false,
	}

	for _, ls := range c.LineScores {
		if ls.IsBatting {
			lineScore = ls
		}
	}

	unparsedOvers := lineScore.Overs
	overs := int(math.Floor(unparsedOvers))
	balls := int(math.Round((unparsedOvers - math.Floor(unparsedOvers)) * 10))

	return domain.TeamScore{
		Team:    domain.NewTeam(c.Team.Name),
		Runs:    domain.NewRuns(lineScore.Runs),
		Wickets: domain.NewWickets(lineScore.Wickets),
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
	Runs      int     `json:"runs"`
	Wickets   int     `json:"wickets"`
	Overs     float64 `json:"overs"`
	IsBatting bool    `json:"isBatting"`
}

type GetScoreResponse struct {
	Header struct {
		Name         string        `json:"name"`
		Competitions []Competition `json:"competitions"`
	} `json:"header"`
	Rosters []Roster `json:"rosters"`
}

type Roster struct {
	HomeAway string          `json:"homeAway"`
	Athletes []RosterAthlete `json:"roster"`
}

type RosterAthlete struct {
	Athlete    Athlete            `json:"athlete"`
	Active     bool               `json:"active"`
	ActiveName string             `json:"activeName"`
	LineScores []AthleteLineScore `json:"linescores"`
}

func (r *RosterAthlete) Facing() bool {
	return r.ActiveName == "striker"
}

type Athlete struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type AthleteLineScore struct {
	Statistics AthleteStatistic `json:"statistics"`
}

type AthleteStatistic struct {
	Categories []AtheleteStatisticCategory `json:"categories"`
}

type AtheleteStatisticCategory struct {
	Name                   string                  `json:"name"`
	AthleteStatisticValues []AthleteStatisticValue `json:"stats"`
}

type AthleteStatisticValue struct {
	Name         string `json:"name"`
	DisplayValue string `json:"displayValue"`
}

func (resp *GetScoreResponse) toScore() (domain.Score, error) {
	var score domain.Score
	var currentState domain.CurrentState

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

	battingTeam := resp.Header.Competitions[0].GetBattingTeam()

	for _, roster := range resp.Rosters {
		if roster.HomeAway != battingTeam.HomeAway {
			continue
		}

		for _, athlete := range roster.Athletes {
			if !athlete.Active {
				continue
			}

			runs, err := athleteStat(athlete, "runs")
			if err != nil {
				return domain.Score{}, err
			}

			balls, err := athleteStat(athlete, "ballsFaced")
			if err != nil {
				return domain.Score{}, err
			}

			batter := domain.Batters{
				CurrentScore: domain.Runs(runs),
				BallsFaced:   domain.Balls(balls),
				Name:         domain.Name(athlete.Athlete.DisplayName),
				Facing:       athlete.Facing(),
			}

			if batter.Facing {
				currentState.BatterA = batter
			} else {
				currentState.BatterB = batter
			}

		}
	}

	for _, roster := range resp.Rosters {
		if roster.HomeAway == battingTeam.HomeAway {
			continue
		}

		for _, athlete := range roster.Athletes {
			if !athlete.Active {
				continue
			}

			wickets, err := athleteStat(athlete, "wickets")
			if err != nil {
				return domain.Score{}, err
			}

			runsConceded, err := athleteStat(athlete, "conceded")
			if err != nil {
				return domain.Score{}, err
			}

			balls, err := athleteStat(athlete, "balls")
			if err != nil {
				return domain.Score{}, err
			}

			overs := balls / 6

			balls = balls % 6

			bowler := domain.Bowler{
				Name:    domain.Name(athlete.Athlete.DisplayName),
				Over:    domain.NewOvers(overs, balls),
				Wickets: domain.NewWickets(wickets),
				Runs:    domain.NewRuns(runsConceded),
			}

			if athlete.ActiveName == "current bowler" {
				currentState.Bowler = bowler
			} else if athlete.ActiveName == "previous bowler" {
				currentState.PrevBowler = bowler
			}
		}
	}

	score.CurrentState = currentState

	return score, nil
}

func athleteStat(athlete RosterAthlete, statname string) (int, error) {
	rawStat := athleteStatString(athlete, statname)

	stat, err := strconv.Atoi(rawStat)
	if err != nil {
		return 0, err
	}
	return stat, nil
}

func athleteStatString(athlete RosterAthlete, statname string) string {
	for _, lineScore := range athlete.LineScores {
		for _, category := range lineScore.Statistics.Categories {
			for _, value := range category.AthleteStatisticValues {
				if value.Name == statname {
					return value.DisplayValue
				}
			}
		}
	}
	return ""
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
