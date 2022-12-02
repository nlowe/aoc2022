package day1

import (
	"fmt"

	"github.com/nlowe/aoc2022/util/gmath"

	"github.com/nlowe/aoc2022/util"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	currentElf := 0
	bestElf := 0
	for line := range challenge.Lines() {
		if line == "" {
			bestElf = gmath.Max(currentElf, bestElf)
			currentElf = 0
			continue
		}

		v := util.MustAtoI(line)
		currentElf += v
	}

	// And once more in case there's no trailing newline
	bestElf = gmath.Max(currentElf, bestElf)
	return bestElf
}
