package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, 0, duplicate_count("abcde"))
}

func Test1(t *testing.T) {
	assert.Equal(t, 1, duplicate_count("abcdea"))
}

func Test2(t *testing.T) {
	assert.Equal(t, 3, duplicate_count("abcdeaB11"))
}

func Test3(t *testing.T) {
	assert.Equal(t, 1, duplicate_count("indivisibility"))
}
