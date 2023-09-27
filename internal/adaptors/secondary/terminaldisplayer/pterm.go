package terminaldisplayer

import (
	"cricwatch/internal/core/domain"
	"errors"

	"github.com/pterm/pterm"
)

type PtermDisplayer struct {
	area *pterm.AreaPrinter
}

func NewPtermDisplayer() *PtermDisplayer {
	return &PtermDisplayer{}
}

func (ptd *PtermDisplayer) Display(score domain.Score) error {

	if ptd.area != nil {
		ptd.area.Stop()
	}

	area, err := pterm.DefaultArea.WithCenter().Start()
	if err != nil {
		return err
	}
	ptd.area = area

	ptd.area.Update(pterm.Sprintf(score.String()))

	return nil
}

func (ptd *PtermDisplayer) Update(score domain.Score) error {
	if ptd.area == nil {
		return errors.New("displayer not initialized")
	}

	ptd.area.Update(pterm.Sprintf(score.String()))

	return nil
}

func (ptd *PtermDisplayer) Close() error {
	if ptd.area == nil {
		return errors.New("displayer not initialized")
	}

	ptd.area.Stop()

	return nil
}
