package day6

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

const partAStartWordLength = 4

func partA(challenge *challenge.Input) int {
	stream := <-challenge.Lines()

	for i := 0; i < len(stream)-partAStartWordLength; i++ {
		word := stream[i : i+partAStartWordLength]

		if isStartWord(word) {
			return i + partAStartWordLength
		}
	}

	panic("no solution")
}

func isStartWord(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				return false
			}
		}
	}

	return true
}
