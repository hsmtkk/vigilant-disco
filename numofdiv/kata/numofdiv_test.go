package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/numofdiv/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n       int
	divisor int
	want    int
}

func TestDivisions(t *testing.T) {
	testCases := []testCase{
		{6, 2, 2},
		{100, 2, 6},
		{2450, 5, 4},
		{9999, 3, 8},
		{2, 3, 0},
		{5, 5, 1},
	}
	for _, tc := range testCases {
		got := kata.Divisions(tc.n, tc.divisor)
		assert.Equal(t, tc.want, got)
	}
}
