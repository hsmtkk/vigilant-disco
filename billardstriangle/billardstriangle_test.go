package billardstriangle_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/billardstriangle"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	testCases := map[int]int{
		1:    1,
		4:    2,
		20:   5,
		100:  13,
		2211: 66,
		9999: 140,
	}
	for input, want := range testCases {
		got := billardstriangle.Solve(input)
		assert.Equal(t, want, got)
	}
}
