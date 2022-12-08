package tilemap

import (
	"fmt"

	"github.com/beefsack/go-astar"

	"github.com/nlowe/aoc2022/challenge"
	"github.com/nlowe/aoc2022/util/gmath"
)

// Container represents the intermediate object stored in the tile map. This container type implements astar.Pather
// with the following assumptions:
//
// * Neighbors on the cardinal direction are reachable if they exist in the tile map
// * The cost to traverse any neighbor is always 1
// * The cost estimate between any two points in the map is the manhattan distance between them
//
// To override these, override Map.NeighborFunc, Map.CostFunc, and Map.EstimateFunc
type Container[T any] struct {
	Value T

	tileMap *Map[T]

	x int
	y int
}

func (t Container[T]) PathNeighbors() (results []astar.Pather) {
	if t.tileMap.NeighborFunc != nil {
		neighbors := t.tileMap.NeighborFunc(t)

		results = make([]astar.Pather, len(neighbors))
		for i, v := range t.tileMap.NeighborFunc(t) {
			results[i] = v
		}

		return
	}

	for _, delta := range []struct {
		x int
		y int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		c, ok := t.tileMap.ContainerAt(t.x+delta.x, t.y+delta.y)
		if !ok {
			continue
		}

		results = append(results, c)
	}

	return
}

func (t Container[T]) PathNeighborCost(to astar.Pather) float64 {
	if t.tileMap.CostFunc != nil {
		return t.tileMap.CostFunc(t, to.(Container[T]))
	}

	// Assume all paths are equal cost
	return 1
}

func (t Container[T]) PathEstimatedCost(to astar.Pather) float64 {
	toSpot := to.(Container[T])

	if t.tileMap.EstimateFunc != nil {
		return t.tileMap.EstimateFunc(t, toSpot)
	}

	return float64(gmath.ManhattanDistance(t.x, t.y, toSpot.x, toSpot.y))
}

// Map represents a fixed size grid of tiles. The top-left tile is [0,0], the bottom-right tile is [w, h]. Tiles are
// stored internally in a Container which implements astar.Pather with the following assumptions:
//
// * Neighbors on the cardinal direction are reachable if they exist in the tile map
// * The cost to traverse any neighbor is always 1
// * The cost estimate between any two points in the map is the manhattan distance between them
//
// To override these, override Map.NeighborFunc, Map.CostFunc, and Map.EstimateFunc
type Map[T any] struct {
	tiles []Container[T]
	w     int
	h     int

	NeighborFunc func(container Container[T]) []Container[T]
	CostFunc     func(a, b Container[T]) float64
	EstimateFunc func(a, b Container[T]) float64
}

// FromInput creates a Map of runes from the input where each line is one row in the map and each rune in each line
// is a column.
func FromInput(input *challenge.Input) *Map[rune] {
	return FromInputOf[rune](input, ToRunes)
}

// FromInputOf creates a Map of the specified type from the input where each line is one row in the map and each
// rune in each line is a column whose value is computed using the provided conversion function.
func FromInputOf[T any](input *challenge.Input, convert func(rune) T) *Map[T] {
	lines := input.LineSlice()

	m := Of[T](len(lines[0]), len(lines))

	for row, line := range lines {
		for column, tile := range line {
			m.SetTile(column, row, convert(tile))
		}
	}

	return m
}

// Of creates a new empty tile map of the specified type and size
func Of[T any](w, h int) *Map[T] {
	return &Map[T]{
		tiles: make([]Container[T], w*h),
		w:     w,
		h:     h,
	}
}

// Size returns the width and height of the tile map
func (t *Map[T]) Size() (int, int) {
	return t.w, t.h
}

func (t *Map[T]) outOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= t.w || y >= t.h
}

func (t *Map[T]) indexOf(x, y int) (int, bool) {
	return x + (t.w * y), !t.outOfBounds(x, y)
}

// SetTile stores the specified tile at location (x,y) in the tile map
func (t *Map[T]) SetTile(x, y int, tile T) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		panic(fmt.Errorf("out of bounds tile access: [%d, %d] is not within the %dx%d map", x, y, t.w, t.h))
	}

	t.tiles[idx] = Container[T]{tileMap: t, Value: tile, x: x, y: y}
}

// ContainerAt returns the tile container at the specified coordinates in the tile map
func (t *Map[T]) ContainerAt(x, y int) (Container[T], bool) {
	idx, ok := t.indexOf(x, y)
	if !ok {
		return Container[T]{}, false
	}

	return t.tiles[idx], true
}

// TileAt returns the value of the Tile at the specified coordinates in the map
func (t *Map[T]) TileAt(x, y int) (T, bool) {
	c, ok := t.ContainerAt(x, y)
	return c.Value, ok
}

// PathBetween uses the A-Star path finding algorithm to find the most efficient path between the two locations in the
// map. By default, the following constraints are used for finding the path:
//
// * Neighbors on the cardinal direction are reachable if they exist in the tile map
// * The cost to traverse any neighbor is always 1
// * The cost estimate between any two points in the map is the manhattan distance between them
//
// To override these, override Map.NeighborFunc, Map.CostFunc, and Map.EstimateFunc
func (t *Map[T]) PathBetween(startX, startY, endX, endY int) ([]astar.Pather, float64, bool) {
	start, ok := t.ContainerAt(startX, startY)
	if !ok {
		return nil, 0, false
	}

	end, ok := t.ContainerAt(endX, endY)
	if !ok {
		return nil, 0, false
	}

	return astar.Path(start, end)
}
