package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/scalesquarestr/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	input := "abcd\nefgh\nijkl\nmnop"
	want := "aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp"
	got := kata.Scale(input, 2, 3)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	input := ""
	want := ""
	got := kata.Scale(input, 5, 5)
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	input := "Kj\nSH"
	want := "Kj\nKj\nSH\nSH"
	got := kata.Scale(input, 1, 2)
	assert.Equal(t, want, got)
}
