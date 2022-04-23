package sqdigseq_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/sqdigseq"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 9, sqdigseq.New().Solve(16))
}

func Test1(t *testing.T) {
	assert.Equal(t, 4, sqdigseq.New().Solve(103))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, sqdigseq.New().Solve(1))
}
