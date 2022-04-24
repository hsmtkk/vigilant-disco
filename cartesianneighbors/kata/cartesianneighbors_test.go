package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/cartesianneighbors/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	got := kata.CartesianNeighbor(2, 2)
	want := [][]int{[]int{1, 1}, []int{1, 2}, []int{1, 3}, []int{2, 1}, []int{2, 3}, []int{3, 1}, []int{3, 2}, []int{3, 3}}
	assert.Equal(t, want, got)
}
