package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/heronformula/kata"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, float32(6.0), kata.Heron(3, 4, 5))
}
