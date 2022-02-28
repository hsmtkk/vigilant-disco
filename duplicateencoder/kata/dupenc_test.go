package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/duplicateencoder/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := "((("
	got := kata.DuplicateEncode("gin")
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := "()()()"
	got := kata.DuplicateEncode("recede")
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := "))(("
	got := kata.DuplicateEncode("(( @")
	assert.Equal(t, want, got)
}

func Test3(t *testing.T) {
	want := ")())())"
	got := kata.DuplicateEncode("Success")
	assert.Equal(t, want, got)
}
