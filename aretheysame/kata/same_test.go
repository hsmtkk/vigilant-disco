package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/aretheysame/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 := []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	assert.True(t, kata.Comp(a1, a2))
}

func Test1(t *testing.T) {
	a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 := []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	assert.False(t, kata.Comp(a1, a2))
}

func Test2(t *testing.T) {
	a2 := []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	assert.False(t, kata.Comp(nil, a2))
}

func Test3(t *testing.T) {
	a1 := []int{2, 2, 3}
	a2 := []int{4, 9, 9}
	assert.True(t, kata.Comp(a1, a2))
}
