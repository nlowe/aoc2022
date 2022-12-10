package day9

import (
	"strings"

	"github.com/nlowe/aoc2022/challenge"
	"github.com/nlowe/aoc2022/util"
	"github.com/nlowe/aoc2022/util/gmath"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

type point struct {
	x int
	y int
}

func (p *point) moveTowards(other point) {
	dx := gmath.Abs(other.x - p.x)
	dy := gmath.Abs(other.y - p.y)

	switch {
	case dx >= 2 && dy == 0:
		if other.x > p.x {
			// R
			p.x++
		} else {
			// L
			p.x--
		}
	case dy >= 2 && dx == 0:
		if other.y > p.y {
			// U
			p.y++
		} else {
			// D
			p.y--
		}
	case dx >= 2 && dy >= 1 || dy >= 2 && dx >= 1:
		// Diagonal
		if other.x > p.x {
			p.x++
			if other.y > p.y {
				p.y++
			} else {
				p.y--
			}
		} else {
			p.x--
			if other.y > p.y {
				p.y++
			} else {
				p.y--
			}
		}
	}
}

func simulate(challenge *challenge.Input, size int) int {
	rope := make([]point, size)

	tailSeen := map[point]struct{}{
		{}: {},
	}

	for move := range challenge.Lines() {
		parts := strings.Fields(move)

		dir := parts[0]
		mag := util.MustAtoI(parts[1])

		for i := 0; i < mag; i++ {
			switch dir {
			case up:
				rope[0].y++
			case right:
				rope[0].x++
			case down:
				rope[0].y--
			case left:
				rope[0].x--
			}

			for n := 1; n < len(rope); n++ {
				rope[n].moveTowards(rope[n-1])
			}

			tailSeen[rope[len(rope)-1]] = struct{}{}
		}
	}

	return len(tailSeen)
}
