package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/lateride/kata"
	"github.com/stretchr/testify/assert"
)

func TestLateRide(t *testing.T) {
	testCases := map[int]int{
		240:  4,
		808:  14,
		1439: 19,
		0:    0,
		23:   5,
		8:    8,
	}
	for input, want := range testCases {
		got := kata.LateRide(input)
		assert.Equal(t, want, got)
	}
}
