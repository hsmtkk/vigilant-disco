package printerror_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hsmtkk/vigilant-disco/printerror"
)

func Test0(t *testing.T) {
	input := "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"
	want := "3/56"
	got := printerror.PrintError(input)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	input := "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"
	want := "6/60"
	got := printerror.PrintError(input)
	assert.Equal(t, want, got)
}
func Test2(t *testing.T) {
	input := "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu"
	want := "11/65"
	got := printerror.PrintError(input)
	assert.Equal(t, want, got)
}
