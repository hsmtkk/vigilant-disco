package hellonameworld_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/hellonameworld"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "Hello, John!", hellonameworld.Hello("johN"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "Hello, Alice!", hellonameworld.Hello("alice"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "Hello, World!", hellonameworld.Hello(""))
}
