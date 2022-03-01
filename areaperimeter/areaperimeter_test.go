package areaperimeter_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/areaperimeter"
	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	const want uint = 32
	got := areaperimeter.NewCalculator().Calculate(6, 10)
	assert.Equal(t, want, got)
}

func TestArea(t *testing.T) {
	const want uint = 9
	got := areaperimeter.NewCalculator().Calculate(3, 3)
	assert.Equal(t, want, got)
}
