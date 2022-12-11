package day11

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 11, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	monkeys := parse(challenge)

	simulateAndSort(monkeys, 20, 0)

	return monkeys[0].inspections * monkeys[1].inspections
}
