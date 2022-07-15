package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/rotmax/kata"
	"github.com/stretchr/testify/assert"
)

func TestRotMax(t *testing.T) {
	testCases := map[int64]int64{
		56789:     68957,
		38458215:  85821534,
		195881031: 988103115,
		896219342: 962193428,
	}
	for input, want := range testCases {
		got := kata.MaxRot(input)
		assert.Equal(t, want, got)
	}
}
