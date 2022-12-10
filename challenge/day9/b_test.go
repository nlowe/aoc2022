package day9

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

const largerExample = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

func TestB(t *testing.T) {
	t.Run("Small", func(t *testing.T) {
		input := challenge.FromLiteral(example)

		result := partB(input)

		require.Equal(t, 1, result)
	})

	t.Run("Large", func(t *testing.T) {
		input := challenge.FromLiteral(largerExample)

		result := partB(input)

		require.Equal(t, 36, result)
	})
}
