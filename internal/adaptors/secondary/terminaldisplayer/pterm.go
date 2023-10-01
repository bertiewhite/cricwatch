package terminaldisplayer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bertiewhite/cricwatch/internal/core/domain"

	"github.com/pterm/pterm"
)

type PtermDisplayer struct {
	area *pterm.AreaPrinter
}

func NewPtermDisplayer() *PtermDisplayer {
	return &PtermDisplayer{}
}

const termWidth = 80

func (ptd *PtermDisplayer) Display(score domain.Score) error {

	if ptd.area != nil {
		ptd.area.Stop()
	}

	area, err := pterm.DefaultArea.Start()
	if err != nil {
		return err
	}
	ptd.area = area

	ptd.Update(score)

	return nil
}

func (ptd *PtermDisplayer) Update(score domain.Score) error {
	if ptd.area == nil {
		return errors.New("displayer not initialized")
	}

	scoreLine := fmt.Sprintf("%s %d-%d (%s) vs %d-%d (%s) %s", score.Home.Team, score.Home.Runs, score.Home.Wickets, score.Home.Overs.String(), score.Away.Runs, score.Away.Wickets, score.Away.Overs.String(), score.Away.Team)

	content := pterm.DefaultBox.
		WithRightPadding(4).
		WithLeftPadding(4).
		WithTitle(scoreLine).
		WithTopPadding(2).
		WithBottomPadding(2).
		WithTitle(scoreLine).
		WithTitleTopCenter().
		Sprintln(createBatterString(score.CurrentState.BatterA, score.CurrentState.BatterB))

	ptd.area.Update(content)

	return nil
}

func (ptd *PtermDisplayer) Close() error {
	if ptd.area == nil {
		return errors.New("displayer not initialized")
	}

	ptd.area.Stop()

	return nil
}

func createBatterString(batterA domain.Batters, batterB domain.Batters) string {
	midWhiteLen := termWidth - len(batterA.Name) - len(batterB.Name)
	midWhite := strings.Repeat(" ", midWhiteLen)
	nameLine := fmt.Sprintf("%s%s%s", batterA.Name, midWhite, batterB.Name)

	scoreA := fmt.Sprintf("%d (%d)", batterA.CurrentScore, batterA.BallsFaced)
	scoreB := fmt.Sprintf("%d (%d)", batterB.CurrentScore, batterB.BallsFaced)
	midWhiteLen = termWidth - len(scoreA) - len(scoreB)
	midWhite = strings.Repeat(" ", midWhiteLen)
	scoreLine := fmt.Sprintf("%s%s%s", scoreA, midWhite, scoreB)

	return fmt.Sprintf("%s\n%s", nameLine, scoreLine)
}

func createBowlerString(currentBowler domain.Bowler, prevBowler domain.Bowler) string {
	return ""

}
