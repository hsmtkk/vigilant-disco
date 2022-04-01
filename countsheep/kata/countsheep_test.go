package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "", countSheep(0))
}

func Test1(t *testing.T) {
	assert.Equal(t, "1 sheep...", countSheep(1))
}

func Test2(t *testing.T) {
	assert.Equal(t, "1 sheep...2 sheep...", countSheep(2))
}
