package day11

import (
	"fmt"

	"github.com/nlowe/aoc2022/util/gmath"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 11, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	monkeys := parse(challenge)

	checks := make([]int, len(monkeys))
	for i, m := range monkeys {
		checks[i] = m.test
	}

	shrink := gmath.LCM(checks...)

	simulateAndSort(monkeys, 10000, shrink)
	return monkeys[0].inspections * monkeys[1].inspections
}
