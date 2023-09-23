package cli

import (
	scoreservice "cricwatch/internal/core/services"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cricwatch",
	Short: "Cricwatch is a CLI tool to get live cricket scores",
	Long:  "Cricwatch is a CLI tool to get live cricket scores",
	Run: func(cmd *cobra.Command, args []string) {

		if svc == nil {
			fmt.Println("Initialization error")
		}

		match, err := svc.SelectMatch()

		if match.Live == false {
			fmt.Println("This match is not live")
			os.Exit(0)
		}

		err = svc.GetAndDisplayScore(match)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

var svc scoreservice.ScoreApplication

func Execute(s scoreservice.ScoreApplication) {

	svc = s

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
