package opseq_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/opseq"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	c := opseq.New()
	assert.Equal(t, 31, c.Calculate([]int{0, 2, 1, -6, -3, 3}))
}
