package correcttimestr_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/correcttimestr"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	times := []string{
		"001122",
		"00;11;22",
		"00:1c:22",
	}
	for _, tm := range times {
		_, err := correcttimestr.Format(tm)
		assert.Error(t, err)
	}
}

func TestFormat(t *testing.T) {
	testCases := map[string]string{
		"09:10:01": "09:10:01",
		"11:70:10": "12:10:10",
		"19:99:09": "20:39:09",
		"19:99:99": "20:40:39",
		"24:01:01": "00:01:01",
		"52:01:01": "04:01:01",
	}
	for input, want := range testCases {
		got, err := correcttimestr.Format(input)
		assert.Nil(t, err)
		assert.Equal(t, want, got)
	}
}
