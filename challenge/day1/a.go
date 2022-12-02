package day1

import (
	"fmt"

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
			if currentElf > bestElf {
				bestElf = currentElf
			}

			currentElf = 0
			continue
		}

		v := util.MustAtoI(line)
		currentElf += v
	}

	if currentElf != 0 {
		if currentElf > bestElf {
			bestElf = currentElf
		}
	}

	return bestElf
}
