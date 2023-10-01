package terminaldisplayer

import (
	"fmt"

	"github.com/bertiewhite/cricwatch/internal/core/domain"
)

type DumbPrinter struct{}

func NewDumbPrinter() *DumbPrinter {
	return &DumbPrinter{}
}

func (dp *DumbPrinter) Display(score domain.Score) error {

	fmt.Println(score.String())

	return nil
}

func (dp *DumbPrinter) Update(score domain.Score) error {

	fmt.Println(score.String())

	return nil
}

func (dp *DumbPrinter) Close() error {
	return nil
}
