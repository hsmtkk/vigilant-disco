package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/twosum/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	got := kata.TwoSum([]int{1, 2, 3}, 4)
	want := [2]int{0, 2}
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	got := kata.TwoSum([]int{1234, 5678, 9012}, 14690)
	want := [2]int{1, 2}
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	got := kata.TwoSum([]int{2, 2, 3}, 4)
	want := [2]int{0, 1}
	assert.Equal(t, want, got)
}
