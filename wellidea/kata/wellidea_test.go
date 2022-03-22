package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/wellidea/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := "I smell a series!"
	got := kata.Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"})
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := "Publish!"
	got := kata.Well([]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"})
	assert.Equal(t, want, got)
}

func Test2(t *testing.T) {
	want := "Fail!"
	got := kata.Well([]string{"bad", "bad"})
	assert.Equal(t, want, got)
}
