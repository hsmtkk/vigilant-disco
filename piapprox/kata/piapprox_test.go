package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/piapprox/kata"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	epsilon   float64
	iteration int
	value     string
}

func TestIterPi(t *testing.T) {
	testCases := []testCase{
		{0.1, 10, "3.0418396189"},
		{0.01, 100, "3.1315929036"},
		{0.001, 1000, "3.1405926538"},
		{0.0001, 10000, "3.1414926536"},
		{0.00001, 100001, "3.1416026535"},
		{0.000001, 1000001, "3.1415936536"},
		{0.05, 20, "3.0916238067"},
	}
	for _, testCase := range testCases {
		iter, calculated := kata.IterPi(testCase.epsilon)
		assert.Equal(t, testCase.iteration, iter)
		assert.Equal(t, testCase.value, calculated)
	}
}
