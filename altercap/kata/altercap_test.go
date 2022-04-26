package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/altercap/kata"
	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	testCases := map[string][]string{
		"abcdef":               []string{"AbCdEf", "aBcDeF"},
		"codewars":             []string{"CoDeWaRs", "cOdEwArS"},
		"abracadabra":          []string{"AbRaCaDaBrA", "aBrAcAdAbRa"},
		"codewarriors":         []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"},
		"indexinglessons":      []string{"InDeXiNgLeSsOnS", "iNdExInGlEsSoNs"},
		"codingisafunactivity": []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"},
	}
	for input, want := range testCases {
		got := kata.Capitalize(input)
		assert.Equal(t, want, got)
	}
}
