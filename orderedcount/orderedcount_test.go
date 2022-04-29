package orderedcount_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/orderedcount"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	want  []orderedcount.Tuple
}

func TestOrderedCount(t *testing.T) {
	testCases := []testCase{
		{"abracadabra", []orderedcount.Tuple{orderedcount.Tuple{'a', 5}, orderedcount.Tuple{'b', 2}, orderedcount.Tuple{'r', 2}, orderedcount.Tuple{'c', 1}, orderedcount.Tuple{'d', 1}}},
		{"Code Wars", []orderedcount.Tuple{orderedcount.Tuple{'C', 1}, orderedcount.Tuple{'o', 1}, orderedcount.Tuple{'d', 1}, orderedcount.Tuple{'e', 1}, orderedcount.Tuple{' ', 1}, orderedcount.Tuple{'W', 1}, orderedcount.Tuple{'a', 1}, orderedcount.Tuple{'r', 1}, orderedcount.Tuple{'s', 1}}},
		{"", []orderedcount.Tuple{}},
	}
	for _, tc := range testCases {
		got := orderedcount.OrderedCount(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
