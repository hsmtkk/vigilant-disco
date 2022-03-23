package parseage_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/parseage"
	"github.com/stretchr/testify/assert"
)

func TestParseAge(t *testing.T) {
	testCases := map[string]uint{
		"2 years old": 2,
		"4 years old": 4,
		"5 years old": 5,
		"7 years old": 7,
	}
	for input, want := range testCases {
		got := parseage.ParseAge(input)
		assert.Equal(t, want, got)
	}
}
