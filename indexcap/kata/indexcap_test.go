package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/indexcap/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	index []int
	want  string
}

func TestCapitalize(t *testing.T) {
	testCases := []testCase{
		{"abcdef", []int{1, 2, 5}, "aBCdeF"},
		{"abcdef", []int{1, 2, 5, 100}, "aBCdeF"},
		{"codewars", []int{1, 3, 5, 50}, "cOdEwArs"},
		{"abracadabra", []int{2, 6, 9, 10}, "abRacaDabRA"},
		{"codewarriors", []int{5}, "codewArriors"},
		{"indexinglessons", []int{0}, "Indexinglessons"},
	}
	for _, tc := range testCases {
		got := kata.Capitalize(tc.input, tc.index)
		assert.Equal(t, tc.want, got)
	}
}
