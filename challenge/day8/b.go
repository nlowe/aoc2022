package day8

import (
	"fmt"

	"github.com/nlowe/aoc2022/util/gmath"

	"github.com/nlowe/aoc2022/util/tilemap"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 8, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	forest := tilemap.FromInputOf(challenge, tilemap.ToInts)

	w, h := forest.Size()

	bestScore := 0

	for x := 1; x < w-1; x++ {
		for y := 1; y < h-1; y++ {
			bestScore = gmath.Max(bestScore, calculateScenicScore(forest, x, y))
		}
	}

	return bestScore
}

func calculateScenicScore(forest *tilemap.Map[int], x, y int) int {
	w, h := forest.Size()

	candidateHeight, _ := forest.TileAt(x, y)

	// north
	northCount := 0
	for dy := y - 1; dy >= 0; dy-- {
		northCount++
		if neighbor, _ := forest.TileAt(x, dy); neighbor >= candidateHeight {
			break
		}
	}

	// south
	southCount := 0
	for dy := y + 1; dy < h; dy++ {
		southCount++

		if neighbor, _ := forest.TileAt(x, dy); neighbor >= candidateHeight {
			break
		}
	}

	// west
	westCount := 0
	for dx := x - 1; dx >= 0; dx-- {
		westCount++

		if neighbor, _ := forest.TileAt(dx, y); neighbor >= candidateHeight {
			break
		}
	}

	// east
	eastCount := 0
	for dx := x + 1; dx < w; dx++ {
		eastCount++
		if neighbor, _ := forest.TileAt(dx, y); neighbor >= candidateHeight {
			break
		}
	}

	return northCount * southCount * westCount * eastCount
}
