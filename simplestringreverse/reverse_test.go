package simplestringreverse_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/simplestringreverse"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testCases := map[string]string{
		"codewars":        "srawedoc",
		"your code":       "edoc ruoy",
		"your code rocks": "skco redo cruoy",
		"i love codewars": "s rawe docevoli",
	}
	for input, want := range testCases {
		got := simplestringreverse.Reverse(input)
		assert.Equal(t, want, got)
	}
}
