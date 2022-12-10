package day9

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 9, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	return simulate(challenge, 10)
}
