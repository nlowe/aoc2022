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

func GCD[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm[T constraints.Integer](a, b T) T {
	return a * b / GCD(a, b)
}

func LCM[T constraints.Integer](vs ...T) T {
	if len(vs) < 2 {
		panic("need at least 2 inputs")
	}

	result := vs[0]
	for i := 1; i < len(vs); i++ {
		result = lcm(result, vs[i])
	}

	return result
}
