package minstep_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/minstep"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 1, minstep.MinimumSteps([]int{4, 6, 3}, 7))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1, minstep.MinimumSteps([]int{10, 9, 9, 8}, 17))
}

func Test2(t *testing.T) {
	assert.Equal(t, 3, minstep.MinimumSteps([]int{8, 9, 10, 4, 2}, 23))
}

func Test3(t *testing.T) {
	assert.Equal(t, 8, minstep.MinimumSteps([]int{19, 98, 69, 28, 75, 45, 17, 98, 67}, 464))
}

func Test4(t *testing.T) {
	assert.Equal(t, 0, minstep.MinimumSteps([]int{4, 6, 3}, 2))
}
