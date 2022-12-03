package day3

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	sum := 0

	unparsed := challenge.LineSlice()
	for i := 0; i < len(unparsed); i += 3 {
		shared := pack(unparsed[i])
		shared &= pack(unparsed[i+2])
		shared &= pack(unparsed[i+1])

		sum += badge(shared)
	}

	return sum
}
