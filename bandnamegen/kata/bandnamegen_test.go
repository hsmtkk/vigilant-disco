package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBandNameGen(t *testing.T) {
	inputWants := map[string]string{
		"knife":   "The Knife",
		"tart":    "Tartart",
		"sandles": "Sandlesandles",
		"bed":     "The Bed",
	}
	for input, want := range inputWants {
		got := bandNameGenerator(input)
		assert.Equal(t, want, got)
	}
}
