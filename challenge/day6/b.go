package day6

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

const partBStartWordLength = 14

func partB(challenge *challenge.Input) int {
	stream := <-challenge.Lines()

	for i := 0; i < len(stream)-partBStartWordLength; i++ {
		word := stream[i : i+partBStartWordLength]

		if isStartWord(word) {
			return i + partBStartWordLength
		}
	}

	panic("no solution")
}
