package keephidrated_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/keephidrated"
	"github.com/stretchr/testify/assert"
)

func TestRequiredWater(t *testing.T) {
	assert.Equal(t, uint(1), keephidrated.RequiredWater(3))
	assert.Equal(t, uint(3), keephidrated.RequiredWater(6.7))
	assert.Equal(t, uint(5), keephidrated.RequiredWater(11.8))
}
