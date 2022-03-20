package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/benefactor/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	input := []float64{14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 16.0}
	assert.Equal(t, int64(628), kata.NewAvg(input, 90))
}

func Test1(t *testing.T) {
	input := []float64{14, 30, 5, 7, 9, 11, 15}
	assert.Equal(t, int64(645), kata.NewAvg(input, 92))
}

func Test2(t *testing.T) {
	input := []float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}
	assert.Equal(t, int64(-1), kata.NewAvg(input, 2000))
}

func Test3(t *testing.T) {
	input := []float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}
	assert.Equal(t, int64(58429), kata.NewAvg(input, 10000))
}
