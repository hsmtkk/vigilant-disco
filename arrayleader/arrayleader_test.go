package arrayleader_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/arrayleader"
	"github.com/stretchr/testify/assert"
)

func TestPositive(t *testing.T) {
	input := []int{1, 2, 3, 4, 0}
	want := []int{4}
	got := arrayleader.ArrayLeader(input)
	assert.Equal(t, want, got)

	input = []int{16, 17, 4, 3, 5, 2}
	want = []int{17, 5, 2}
	got = arrayleader.ArrayLeader(input)
	assert.Equal(t, want, got)
}

func TestNegative(t *testing.T) {
	input := []int{-1, -29, -26, -2}
	want := []int{-1}
	got := arrayleader.ArrayLeader(input)
	assert.Equal(t, want, got)

	input = []int{-36, -12, -27}
	want = []int{-36, -12}
	got = arrayleader.ArrayLeader(input)
	assert.Equal(t, want, got)
}

func TestMix(t *testing.T) {
	input := []int{5, -2, 2}
	want := []int{5, 2}
	got := arrayleader.ArrayLeader(input)
	assert.Equal(t, want, got)

	input = []int{0, -1, -29, 3, 2}
	want = []int{0, -1, 3, 2}
	got = arrayleader.ArrayLeader(input)
	assert.Equal(t, want, got)
}
