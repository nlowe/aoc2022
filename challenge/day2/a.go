package day2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

const (
	theyPlayRock     = "A"
	theyPlayPaper    = "B"
	theyPlayScissors = "C"

	youPlayRock     = "X"
	youPlayPaper    = "Y"
	youPlayScissors = "Z"

	outcomeLoss = 0
	outcomeTie  = 3
	outcomeWin  = 6
)

var exactScore = map[string]map[string]int{
	theyPlayRock: {
		youPlayRock:     outcomeTie + valueOf(youPlayRock),
		youPlayPaper:    outcomeWin + valueOf(youPlayPaper),
		youPlayScissors: outcomeLoss + valueOf(youPlayScissors),
	},
	theyPlayPaper: {
		youPlayRock:     outcomeLoss + valueOf(youPlayRock),
		youPlayPaper:    outcomeTie + valueOf(youPlayPaper),
		youPlayScissors: outcomeWin + valueOf(youPlayScissors),
	},
	theyPlayScissors: {
		youPlayRock:     outcomeWin + valueOf(youPlayRock),
		youPlayPaper:    outcomeLoss + valueOf(youPlayPaper),
		youPlayScissors: outcomeTie + valueOf(youPlayScissors),
	},
}

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	totalScore := 0
	for round := range challenge.Lines() {
		selection := strings.Fields(round)
		totalScore += exactScore[selection[0]][selection[1]]
	}

	return totalScore
}

func valueOf(play string) int {
	return 1 + int(play[0]-youPlayRock[0])
}
