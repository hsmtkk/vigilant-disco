package multiplicationtable_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/multiplicationtable"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	want := [][]int{{1}}
	got := multiplicationtable.MultiplicationTable(1)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := [][]int{{1, 2}, {2, 4}}
	got := multiplicationtable.MultiplicationTable(2)
	assert.Equal(t, want, got)
}

func Test3(t *testing.T) {
	want := [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}
	got := multiplicationtable.MultiplicationTable(3)
	assert.Equal(t, want, got)
}
