package day2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

const (
	desireLoss = "X"
	desireTie  = "Y"
	desireWin  = "Z"
)

var forDesiredOutcome = map[string]map[string]string{
	theyPlayRock: {
		desireLoss: youPlayScissors,
		desireTie:  youPlayRock,
		desireWin:  youPlayPaper,
	},
	theyPlayPaper: {
		desireLoss: youPlayRock,
		desireTie:  youPlayPaper,
		desireWin:  youPlayScissors,
	},
	theyPlayScissors: {
		desireLoss: youPlayPaper,
		desireTie:  youPlayScissors,
		desireWin:  youPlayRock,
	},
}

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	totalScore := 0
	for round := range challenge.Lines() {
		selection := strings.Fields(round)
		totalScore += exactScore[selection[0]][forDesiredOutcome[selection[0]][selection[1]]]
	}

	return totalScore
}
