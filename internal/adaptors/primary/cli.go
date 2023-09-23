package cli

import (
	"cricwatch/internal/core/domain"
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

		err := svc.GetAndDisplayScore(domain.Match{ID: 1336129})
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
