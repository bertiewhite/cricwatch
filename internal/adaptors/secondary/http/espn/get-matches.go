package espn

import (
	"cricwatch/internal/core/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const matchesEndpoint = "/scorepanel"

type League struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status struct {
		Type struct {
			Detail string `json:"detail"`
		} `json:"type"`
	} `json:"status"`
}

func (e Event) isLive() bool {
	return e.Status.Type.Detail == "Live"
}

type Score struct {
	Leagues []League `json:"leagues"`
	Events  []Event  `json:"events"`
}

type GetMatchesResponse struct {
	Scores []Score `json:"scores"`
}

func (resp *GetMatchesResponse) toMatches() domain.Matches {
	matches := make(domain.Matches)

	for _, score := range resp.Scores {
		league := domain.League("")
		// ESPN Api only seems to return one league at a time but if there's multiple grab the first
		// could be a bug where weird leagues get reported but thats fine with me
		if len(score.Leagues) >= 1 {
			league = domain.League(score.Leagues[0].Name)
		}

		for _, event := range score.Events {

			id, err := strconv.Atoi(event.ID)
			if err != nil {
				panic("oh god no please this wasn't meant to happen I was so confident")
			}

			match := domain.Match{
				ID:     domain.MatchID(id),
				Name:   domain.Name(event.Name),
				League: league,
				Live:   event.isLive(),
			}

			matches[league] = append(matches[league], match)
		}
	}

	return matches
}

func (s *Client) GetMatches() (domain.Matches, error) {
	url, err := url.JoinPath(baseUrl, matchesEndpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ESPN API returned non-200 status code: %d", resp.StatusCode)
	}

	var getMatchesResponse GetMatchesResponse
	err = json.NewDecoder(resp.Body).Decode(&getMatchesResponse)
	if err != nil {
		return nil, err
	}

	matches := getMatchesResponse.toMatches()

	return matches, nil
}
