package day3

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	sum := 0
	for rucksack := range challenge.Lines() {
		left, right := partialPack(rucksack)

		sum += badge(left & right)
	}

	return sum
}

func partialPack(rucksack string) (left, right uint64) {
	return pack(rucksack[:len(rucksack)/2]), pack(rucksack[len(rucksack)/2:])
}

func pack(rucksack string) uint64 {
	var kinds uint64
	for _, item := range rucksack {
		if item >= 'a' && item <= 'z' {
			kinds |= 1 << (item - 'a')
		} else {
			kinds |= 1 << (item - 'A' + 26)
		}
	}

	return kinds
}

func badge(rucksack uint64) int {
	var sum int

	for i := 0; i < 26*2; i++ {
		if (rucksack>>i)&0b1 == 1 {
			sum += i + 1
		}
	}

	return sum
}
