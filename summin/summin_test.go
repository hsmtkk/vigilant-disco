package summin_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/summin"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	input := [][]uint{
		{7, 9, 8, 6}, {6, 5, 4, 3}, {5, 7, 4, 5}, {7, 9, 4, 3},
	}
	const want uint = 16
	got := summin.SumMin(input)
	assert.Equal(t, want, got)
}
