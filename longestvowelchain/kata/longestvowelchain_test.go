package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/longestvowelchain/kata"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	inputWants := map[string]int{
		"codewarriors":          2,
		"suoidea":               3,
		"ultrarevolutionariees": 3,
		"strengthlessnesses":    1,
		"cuboideonavicuare":     2,
		"chrononhotonthuooaos":  5,
		"iiihoovaeaaaoougjyaw":  8,
	}
	for input, want := range inputWants {
		got := kata.Solve(input)
		assert.Equal(t, want, got)
	}
}
