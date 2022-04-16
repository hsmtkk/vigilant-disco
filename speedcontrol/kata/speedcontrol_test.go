package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/speedcontrol/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 41, kata.Gps(20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 219, kata.Gps(12, []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}))
}
