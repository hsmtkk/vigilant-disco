package kata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hsmtkk/vigilant-disco/mumbling/kata"
)

func TestAccum(t *testing.T) {
	inputWant := map[string]string{
		"ZpglnRxqenU": "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu",
		"NyffsGeyylB": "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb",
		"MjtkuBovqrU": "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu",
		"EvidjUnokmM": "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm",
		"HbideVbxncC": "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc",
	}
	for input, want := range inputWant {
		got := kata.Accum(input)
		assert.Equal(t, want, got)
	}
}
