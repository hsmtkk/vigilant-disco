package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/maxlendiff/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	s1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	s2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	want := 13
	got := kata.MxDifLg(s1, s2)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	s1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 := []string{"bbbaaayddqbbrrrv"}
	want := 10
	got := kata.MxDifLg(s1, s2)
	assert.Equal(t, want, got)
}

func TestXxx(t *testing.T) {
	s1 := []string{}
	s2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	want := -1
	got := kata.MxDifLg(s1, s2)
	assert.Equal(t, want, got)
}
