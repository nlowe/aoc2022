package day8

import (
	"fmt"

	"github.com/nlowe/aoc2022/util/tilemap"

	"github.com/spf13/cobra"

	"github.com/nlowe/aoc2022/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 8, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	forest := tilemap.FromInputOf(challenge, tilemap.ToInts)

	w, h := forest.Size()

	// The perimeter is always visible
	visible := 2*w + (h-2)*2

	for x := 1; x < w-1; x++ {
		for y := 1; y < h-1; y++ {
			if isVisible(forest, x, y) {
				visible++
			}
		}
	}

	return visible
}

func isVisible(forest *tilemap.Map[int], x, y int) bool {
	w, h := forest.Size()

	treeHeight, _ := forest.TileAt(x, y)

	// north
	visibleNorth := true
	for dy := y - 1; dy >= 0; dy-- {
		if neighbor, _ := forest.TileAt(x, dy); neighbor >= treeHeight {
			visibleNorth = false
			break
		}
	}

	// south
	visibleSouth := true
	for dy := y + 1; dy <= h; dy++ {
		if neighbor, _ := forest.TileAt(x, dy); neighbor >= treeHeight {
			visibleSouth = false
			break
		}
	}

	// west
	visibleEast := true
	for dx := x - 1; dx >= 0; dx-- {
		if neighbor, _ := forest.TileAt(dx, y); neighbor >= treeHeight {
			visibleEast = false
			break
		}
	}

	// east
	visibleWest := true
	for dx := x + 1; dx <= w; dx++ {
		if neighbor, _ := forest.TileAt(dx, y); neighbor >= treeHeight {
			visibleWest = false
			break
		}
	}

	return visibleNorth || visibleSouth || visibleEast || visibleWest
}
