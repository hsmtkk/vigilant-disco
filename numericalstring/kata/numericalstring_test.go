package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/numericalstring/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input string
	want  string
}

func TestNumericals(t *testing.T) {
	testCases := []testCase{
		{"Hello, World!", "1112111121311"},
		{"Inconceivable!", "11112211111121"},
		{"hello hello", "11121122342"},
		{"Hello", "11121"},
		{"aaaaaaaaaaaa", "123456789101112"},
	}
	for _, tc := range testCases {
		got := kata.Numericals(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
