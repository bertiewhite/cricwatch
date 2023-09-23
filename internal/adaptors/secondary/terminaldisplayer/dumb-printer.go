package terminaldisplayer

import (
	"cricwatch/internal/core/domain"
	"fmt"
)

type DumbPrinter struct{}

func NewDumbPrinter() *DumbPrinter {
	return &DumbPrinter{}
}

func (dp *DumbPrinter) Display(score domain.Score) error {

	fmt.Println(score.String())

	return nil
}
