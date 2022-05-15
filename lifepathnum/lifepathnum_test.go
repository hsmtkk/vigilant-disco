package lifepathnum_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/lifepathnum"
	"github.com/stretchr/testify/assert"
)

func TestLifePathNumber(t *testing.T) {
	testCases := map[string]int{
		"1879-03-14": 6,
		"1815-12-10": 1,
		"1961-07-04": 1,
		"1955-10-28": 4,
		"1452-04-15": 4,
		"1791-12-26": 2,
		"1906-12-09": 1,
		"1912-06-23": 6,
		"1950-08-11": 7,
		"1956-01-31": 8,
		"1965-04-14": 3,
		"1971-06-28": 7,
	}
	for input, want := range testCases {
		got := lifepathnum.LifePathNumber(input)
		assert.Equal(t, want, got)
	}
}
