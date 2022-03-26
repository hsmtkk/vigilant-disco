package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/sumintstring/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 16, kata.SumOfIntegersInString("12.4"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 3, kata.SumOfIntegersInString("h3ll0w0rld"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 3635, kata.SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
}

func TestParse(t *testing.T) {
	input := "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"
	want := []string{"30", "20", "10", "0", "1203", "914", "3", "1349", "102", "4"}
	got := kata.Parse(input)
	assert.Equal(t, want, got)
}
