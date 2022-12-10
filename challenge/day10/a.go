package day10

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
		Short: "Day 10, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	sum := 0
	cycles := 0
	x := 1
	for instr := range challenge.Lines() {
		cycles++
		sum += cycleSum(cycles, x)

		parts := strings.Fields(instr)

		switch parts[0] {
		case "noop":
		case "addx":
			cycles++
			sum += cycleSum(cycles, x)
			x += util.MustAtoI(parts[1])
		}
	}

	return sum
}

func cycleSum(cycles int, x int) int {
	if cycles == 20 || (cycles-20)%40 == 0 {
		return cycles * x
	}

	return 0
}
