package cli

import (
	"cricwatch/internal/adaptors/secondary/http/espn"
	"cricwatch/internal/adaptors/secondary/terminaldisplayer"
	scoreservice "cricwatch/internal/core/services"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cricwatch",
	Short: "Cricwatch is a CLI tool to get live cricket scores",
	Long:  "Cricwatch is a CLI tool to get live cricket scores",
	Run: func(cmd *cobra.Command, args []string) {
		espnClient := espn.NewEspnClient(http.DefaultClient)
		dumDisply := terminaldisplayer.NewDumbPrinter()
		svc := scoreservice.NewScoreService(espnClient, dumDisply)
		err := svc.GetAndDisplayScore(1336129)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
