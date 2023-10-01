package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bertiewhite/cricwatch/internal/core/domain"
	scoreservice "github.com/bertiewhite/cricwatch/internal/core/services"

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

		matchID, err := cmd.Flags().GetString("matchid")
		if err != nil {
			fmt.Println(fmt.Sprintf("Error while parsing matchid flag: %s", err.Error()))
		}

		var match domain.Match
		if matchID != "" {
			id, err := strconv.Atoi(matchID)
			if err != nil {
				fmt.Println("Invalid match id")
				os.Exit(1)
			}
			match = domain.Match{
				ID: domain.MatchID(id),
			}
		} else {
			tmpmatch, err := svc.SelectMatch()
			match = tmpmatch
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println(fmt.Sprintf("Temp Debug: selected match id: %d", match.ID.Int()))
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

	rootCmd.Flags().String("matchid", "", "Specify match id to skip match selection")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
