package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/makevalley/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, []int{17, 15, 8, 7, 4, 1, 4, 5, 7, 14, 17}, kata.MakeValley([]int{17, 17, 15, 14, 8, 7, 7, 5, 4, 4, 1}))
}

func Test1(t *testing.T) {
	assert.Equal(t, []int{20, 6, 2, 7}, kata.MakeValley([]int{20, 7, 6, 2}))
}

func Test2(t *testing.T) {
	assert.Equal(t, []int{14, 8, 10}, kata.MakeValley([]int{14, 10, 8}))
}

func Test3(t *testing.T) {
	assert.Equal(t, []int{20, 17, 12, 10, 4, 2, 1, 1, 2, 9, 12, 13, 18}, kata.MakeValley([]int{20, 18, 17, 13, 12, 12, 10, 9, 4, 2, 2, 1, 1}))
}
