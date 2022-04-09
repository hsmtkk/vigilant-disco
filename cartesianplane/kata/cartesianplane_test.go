package kata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/hsmtkk/vigilant-disco/cartesianplane/kata"
)

func TestSumin(t *testing.T) {
	assert.Equal(t, int64(55), kata.SuMin(5))
	assert.Equal(t, int64(338350), kata.SuMin(100))
}

func TestSumax(t *testing.T) {
	assert.Equal(t, int64(95), kata.SuMax(5))
	assert.Equal(t, int64(671650), kata.SuMax(100))
}
