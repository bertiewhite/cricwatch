package main

import (
	cli "cricwatch/internal/adaptors/primary"
	"cricwatch/internal/adaptors/secondary/http/espn"
	"cricwatch/internal/adaptors/secondary/terminaldisplayer"
	scoreservice "cricwatch/internal/core/services"
	"net/http"
)

func main() {

	espnClient := espn.NewClient(http.DefaultClient)
	dumPrinter := terminaldisplayer.NewDumbPrinter()

	svc := scoreservice.NewScoreService(espnClient, dumPrinter)

	cli.Execute(svc)
}
