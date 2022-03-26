package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/prevmulthree/kata"
	"github.com/stretchr/testify/assert"
)

func TestPrevMultOfThree(t *testing.T) {
	inputWant := map[int]interface{}{
		1:      nil,
		25:     nil,
		36:     36,
		1244:   12,
		952406: 9,
	}
	for input, want := range inputWant {
		got := kata.PrevMultOfThree(input)
		assert.Equal(t, want, got)
	}
}
