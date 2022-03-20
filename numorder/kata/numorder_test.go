package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/numorder/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.True(t, kata.InAscOrder([]int{1, 2, 4, 7, 19}))
}

func Test1(t *testing.T) {
	assert.True(t, kata.InAscOrder([]int{1, 2, 3, 4, 5}))
}

func Test2(t *testing.T) {
	assert.False(t, kata.InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}))
}

func Test3(t *testing.T) {
	assert.False(t, kata.InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
