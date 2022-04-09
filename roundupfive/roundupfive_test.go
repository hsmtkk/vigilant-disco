package roundupfive_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/roundupfive"
	"github.com/stretchr/testify/assert"
)

func TestRoundUpFive(t *testing.T) {
	inputWants := map[int]int{
		0:  0,
		2:  5,
		3:  5,
		12: 15,
		21: 25,
		30: 30,
		-2: 0,
		-5: -5,
	}
	for input, want := range inputWants {
		got := roundupfive.RoundUpFive(input)
		assert.Equal(t, want, got)
	}
}
