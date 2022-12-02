package day1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

const testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestA(t *testing.T) {
	input := challenge.FromLiteral(testInput)

	result := partA(input)

	require.Equal(t, 24000, result)
}
