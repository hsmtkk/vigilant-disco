package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/allunique/kata"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	testCases := map[string]bool{
		"  nAa": false,
		"abcde": true,
		"++-":   false,
		"AaBbC": true,
	}
	for input, want := range testCases {
		got := kata.HasUniqueChar(input)
		assert.Equal(t, want, got)
	}
}
