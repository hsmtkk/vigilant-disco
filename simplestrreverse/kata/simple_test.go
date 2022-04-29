package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/simplestrreverse/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	begin int
	end   int
	want  string
}

func TestSolve(t *testing.T) {
	testCases := []testCase{
		{"codewars", 1, 5, "cawedors"},
		{"codingIsFun", 2, 100, "conuFsIgnid"},
		{"FunctionalProgramming", 2, 15, "FuargorPlanoitcnmming"},
		{"abcefghijklmnopqrstuvwxyz", 0, 20, "vutsrqponmlkjihgfecbawxyz"},
		{"abcefghijklmnopqrstuvwxyz", 5, 20, "abcefvutsrqponmlkjihgwxyz"},
	}
	for _, tc := range testCases {
		got := kata.Solve(tc.input, tc.begin, tc.end)
		assert.Equal(t, tc.want, got)
	}
}
