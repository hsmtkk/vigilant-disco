package kata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hsmtkk/vigilant-disco/closestten/kata"
)

func TestXxx(t *testing.T) {
	assert.Equal(t, uint32(50), kata.ClosestMultipleOf10(54))
	assert.Equal(t, uint32(60), kata.ClosestMultipleOf10(55))
}
