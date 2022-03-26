package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/fusc/kata"
	"github.com/stretchr/testify/assert"
)

func TestFusc(t *testing.T) {
	assert.Equal(t, 21, kata.Fusc(85))
}
