package freqseq_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/freqseq"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	orig      string
	separator string
	want      string
}

func TestFreqSeq(t *testing.T) {
	testCases := []testCase{
		{"hello world", "-", "1-1-3-3-2-1-1-2-1-3-1"},
		{"19999999", ":", "1:7:7:7:7:7:7:7"},
		{"^^^**$", "x", "3x3x3x2x2x1"},
	}
	for _, tc := range testCases {
		got := freqseq.FreqSeq(tc.orig, tc.separator)
		assert.Equal(t, tc.want, got)
	}
}
