package main

import (
	cli "cricwatch/internal/adaptors/primary"
	"cricwatch/internal/adaptors/secondary/http/espn"
	"cricwatch/internal/adaptors/secondary/terminaldisplayer"
	"cricwatch/internal/adaptors/secondary/userinput/pterminput"
	scoreservice "cricwatch/internal/core/services"
	"net/http"
)

func main() {

	espnClient := espn.NewClient(http.DefaultClient)
	dumPrinter := terminaldisplayer.NewDumbPrinter()
	ptermInputter := pterminput.New()

	svc := scoreservice.NewScoreService(espnClient, dumPrinter, ptermInputter)

	cli.Execute(svc)
}
