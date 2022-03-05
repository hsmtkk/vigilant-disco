package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/targetdate/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	got := kata.DateNbDays(4281, 5087, 2)
	want := "2024-07-03"
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	got := kata.DateNbDays(4620, 5188, 2)
	want := "2021-09-19"
	assert.Equal(t, want, got)
}
