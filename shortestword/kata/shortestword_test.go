package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/shortestword/kata"
	"github.com/stretchr/testify/assert"
)

func TestFindShort(t *testing.T) {
	inputWant := map[string]int{
		"bitcoin take over the world maybe who knows perhaps":                3,
		"turns out random test cases are easier than writing out basic ones": 3,
		"Let's travel abroad shall we":                                       2,
	}
	for input, want := range inputWant {
		got := kata.FindShort(input)
		assert.Equal(t, want, got)
	}
}
