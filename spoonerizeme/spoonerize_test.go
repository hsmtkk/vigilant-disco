package spoonerizeme_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/spoonerizeme"
	"github.com/stretchr/testify/assert"
)

func TestSpoonerize(t *testing.T) {
	inputWant := map[string]string{
		"nit picking":   "pit nicking",
		"wedding bells": "bedding wells",
		"jelly beans":   "belly jeans",
	}
	for input, want := range inputWant {
		got := spoonerizeme.Spoonerize(input)
		assert.Equal(t, want, got)
	}
}
