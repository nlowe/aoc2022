package day4

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	count := 0
	for pairs := range challenge.Lines() {
		assignments := strings.Split(pairs, ",")

		aStart, aEnd := parse(assignments[0])
		bStart, bEnd := parse(assignments[1])

		if overlaps(aStart, aEnd, bStart, bEnd) {
			count++
		}
	}

	return count
}

func overlaps(aStart, aEnd, bStart, bEnd int) bool {
	return (bStart <= aStart && bEnd >= aStart) ||
		(bStart <= aEnd && bEnd >= aEnd) ||
		(aStart <= bStart && aEnd >= bEnd) ||
		(aStart <= bEnd && aEnd >= bEnd)
}
