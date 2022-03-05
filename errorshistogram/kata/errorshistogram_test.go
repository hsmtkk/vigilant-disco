package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/errorshistogram/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	input := "tpwaemuqxdmwqbqrjbeosjnejqorxdozsxnrgpgqkeihqwybzyymqeazfkyiucesxwutgszbenzvgxibxrlvmzihcb"
	want := "u  3     ***\rw  4     ****\rx  6     ******\rz  6     ******"
	got := kata.Hist(input)
	assert.Equal(t, want, got)
}

func TestCountChar(t *testing.T) {
	input := "tpwaemuqxdmwqbqrjbeosjnejqorxdozsxnrgpgqkeihqwybzyymqeazfkyiucesxwutgszbenzvgxibxrlvmzihcb"
	want := map[rune]uint{'u': 3, 'w': 4, 'x': 6, 'z': 6}
	got := kata.CountChar(input)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	input := "aaifzlnderpeurcuqjqeywdq"
	want := "u  2     **\rw  1     *\rz  1     *"
	got := kata.Hist(input)
	assert.Equal(t, want, got)
}
