package doubleton_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/doubleton"
	"github.com/stretchr/testify/assert"
)

func TestNextDoubleton(t *testing.T) {
	assert.Equal(t, 10, doubleton.NextDoubleton(1))
	assert.Equal(t, 12, doubleton.NextDoubleton(10))
	assert.Equal(t, 121, doubleton.NextDoubleton(120))
	assert.Equal(t, 1311, doubleton.NextDoubleton(1234))
}
