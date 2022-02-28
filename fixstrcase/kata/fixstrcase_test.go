package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/fixstrcase/kata"
	"github.com/stretchr/testify/assert"
)

type InputWant struct {
	input string
	want  string
}

func Test0(t *testing.T) {
	inputWants := []InputWant{
		{input: "a", want: "a"},
		{input: "Z", want: "Z"},
		{input: "coDe", want: "code"},
		{input: "CODe", want: "CODE"},
		{input: "coDE", want: "code"},
		{input: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", want: "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
	}
	for _, iw := range inputWants {
		got := kata.Solve(iw.input)
		assert.Equal(t, iw.want, got)
	}
}
