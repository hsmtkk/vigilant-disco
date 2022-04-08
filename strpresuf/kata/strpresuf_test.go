package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/strpresuf/kata"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	inputWant := map[string]int{
		"abcd":     0,
		"abcda":    1,
		"abcdabc":  3,
		"abcabc":   3,
		"abcabca":  1,
		"abcdabcc": 0,
		"aaaaa":    2,
		"aaaa":     2,
		"aaa":      1,
		"aa":       1,
		"a":        0,
		"acbacc":   0,
	}
	for input, want := range inputWant {
		got := kata.Solve(input)
		assert.Equal(t, want, got)
	}
}
