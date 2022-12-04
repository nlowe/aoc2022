package day4

import (
	"fmt"
	"strings"

	"github.com/nlowe/aoc2022/util"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	count := 0
	for pairs := range challenge.Lines() {
		assignments := strings.Split(pairs, ",")

		aStart, aEnd := parse(assignments[0])
		bStart, bEnd := parse(assignments[1])

		if contains(aStart, aEnd, bStart, bEnd) {
			count++
		}
	}

	return count
}

func contains(aStart, aEnd, bStart, bEnd int) bool {
	return (bStart >= aStart && bEnd <= aEnd) ||
		(aStart >= bStart && aEnd <= bEnd)
}

func parse(assignment string) (start, end int) {
	parts := strings.Split(assignment, "-")
	return util.MustAtoI(parts[0]), util.MustAtoI(parts[1])
}
