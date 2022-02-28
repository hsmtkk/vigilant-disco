package secminhour_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/secminhour"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	inputWants := map[uint]string{
		3600:   "1 hour(s) and 0 minute(s)",
		3601:   "1 hour(s) and 0 minute(s)",
		3500:   "0 hour(s) and 58 minute(s)",
		323500: "89 hour(s) and 51 minute(s)",
		0:      "0 hour(s) and 0 minute(s)",
	}
	for input, want := range inputWants {
		got := secminhour.ToTime(input)
		assert.Equal(t, want, got)
	}
}
