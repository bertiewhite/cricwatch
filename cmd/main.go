package main

import (
	"net/http"

	cli "github.com/bertiewhite/cricwatch/internal/adaptors/primary"
	"github.com/bertiewhite/cricwatch/internal/adaptors/secondary/http/espn"
	"github.com/bertiewhite/cricwatch/internal/adaptors/secondary/terminaldisplayer"
	"github.com/bertiewhite/cricwatch/internal/adaptors/secondary/userinput/pterminput"
	scoreservice "github.com/bertiewhite/cricwatch/internal/core/services"
)

func main() {

	espnClient := espn.NewClient(http.DefaultClient)
	// dumPrinter := terminaldisplayer.NewDumbPrinter()
	ptermOutput := terminaldisplayer.NewPtermDisplayer()

	ptermInputter := pterminput.New()

	svc := scoreservice.NewScoreService(espnClient, ptermOutput, ptermInputter)

	cli.Execute(svc)
}
