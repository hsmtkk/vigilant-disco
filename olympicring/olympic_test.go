package olympicring_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/olympicring"
	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	testCases := map[string]string{
		"wHjMudLwtoPGocnJ":     "Bronze!",
		"eCEHWEPwwnvzMicyaRjk": "Bronze!",
		"JKniLfLW":             "Not even a medal!",
		"EWlZlDFsEIBufsalqof":  "Silver!",
		"IMBAWejlGRTDWetPS":    "Gold!",
	}
	for input, want := range testCases {
		got := olympicring.Score(input)
		assert.Equal(t, want, got)
	}
}
