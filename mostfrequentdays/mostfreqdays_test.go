package mostfrequentdays_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/mostfrequentdays"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	yearWant := map[int][]string{
		2427: {"Friday"},
		2185: {"Saturday"},
		1770: {"Monday"},
		1785: {"Saturday"},
		1984: {"Monday", "Sunday"},
		3076: {"Saturday", "Sunday"},
	}
	for year, want := range yearWant {
		got := mostfrequentdays.MostFrequentDays(year)
		assert.Equal(t, want, got)
	}
}
