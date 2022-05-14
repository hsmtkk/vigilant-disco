package stringdoubles_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/stringdoubles"
	"github.com/stretchr/testify/assert"
)

func TestShrink(t *testing.T) {
	testCases := map[string]string{
		"abbbzz":          "ab",
		"zzzzykkkd":       "ykd",
		"abbcccdddda":     "aca",
		"vvvvvoiiiiin":    "voin",
		"rrrmooomqqqqj":   "rmomj",
		"xxbnnnnnyaaaaam": "bnyam",
	}
	for input, want := range testCases {
		got := stringdoubles.Shrink(input)
		assert.Equal(t, want, got)
	}
}
