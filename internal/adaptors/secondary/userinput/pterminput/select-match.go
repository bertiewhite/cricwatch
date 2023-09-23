package pterminput

import (
	"cricwatch/internal/core/domain"
	"errors"

	"github.com/pterm/pterm"
)

func (p *PTermInput) SelectMatch(matches domain.Matches) (domain.Match, error) {
	var leagueOptions []string

	for league, _ := range matches {
		leagueOptions = append(leagueOptions, league.String())
	}

	selectedLeague, err := pterm.DefaultInteractiveSelect.WithOptions(leagueOptions).Show()

	if err != nil {
		return domain.Match{}, err
	}

	league := domain.League(selectedLeague)
	selectedMatches := matches[league]

	var matchOptions []string
	for _, match := range selectedMatches {
		matchOptions = append(matchOptions, match.Name.String())
	}
	selectedMatch, err := pterm.DefaultInteractiveSelect.WithOptions(matchOptions).Show()
	if err != nil {
		return domain.Match{}, err
	}

	for _, match := range selectedMatches {
		if match.Name.String() == selectedMatch {
			return match, nil
		}
	}

	return domain.Match{}, errors.New("errrrrr how the hell did this happen")
}
