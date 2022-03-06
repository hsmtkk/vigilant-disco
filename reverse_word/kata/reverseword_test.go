package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/reverse_word/kata"
	"github.com/stretchr/testify/assert"
)

type inputWant struct {
	input string
	want  string
}

func TestReverseWords(t *testing.T) {
	inputWants := []inputWant{
		{"The quick brown fox jumps over the lazy dog.", "ehT kciuq nworb xof spmuj revo eht yzal .god"},
		{"apple", "elppa"},
		{"a b c d", "a b c d"},
		{"double  spaced  words", "elbuod  decaps  sdrow"},
		{"stressed desserts", "desserts stressed"},
		{"hello hello", "olleh olleh"},
	}
	for _, iw := range inputWants {
		got := kata.ReverseWords(iw.input)
		assert.Equal(t, iw.want, got)
	}
}
