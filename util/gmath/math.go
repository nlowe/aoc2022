package gmath

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Clamp[T Number](low, n, high T) T {
	if low > high {
		panic(fmt.Errorf("Clamp: low cannot be > high: %v > %v", low, high))
	}

	return Max(Min(n, high), low)
}

func ManhattanDistance[T Number](x1, y1, x2, y2 T) T {
	return Abs(x2-x1) + Abs(y2-y1)
}
