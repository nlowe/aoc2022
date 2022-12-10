package day10

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nlowe/aoc2022/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(example)

	result := partB(input)

	require.Equal(t, `
██⠀⠀██⠀⠀██⠀⠀██⠀⠀██⠀⠀██⠀⠀██⠀⠀██⠀⠀██⠀⠀██⠀⠀
███⠀⠀⠀███⠀⠀⠀███⠀⠀⠀███⠀⠀⠀███⠀⠀⠀███⠀⠀⠀███⠀
████⠀⠀⠀⠀████⠀⠀⠀⠀████⠀⠀⠀⠀████⠀⠀⠀⠀████⠀⠀⠀⠀
█████⠀⠀⠀⠀⠀█████⠀⠀⠀⠀⠀█████⠀⠀⠀⠀⠀█████⠀⠀⠀⠀⠀
██████⠀⠀⠀⠀⠀⠀██████⠀⠀⠀⠀⠀⠀██████⠀⠀⠀⠀⠀⠀████
███████⠀⠀⠀⠀⠀⠀⠀███████⠀⠀⠀⠀⠀⠀⠀███████⠀⠀⠀⠀⠀
`, result)
}
