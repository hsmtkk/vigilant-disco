package countip_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/countip"
	"github.com/stretchr/testify/assert"
)

type IPCounter interface {
	CountIP(from, to string) int
}

func Test0(t *testing.T) {
	var cnt IPCounter = countip.New()
	assert.Equal(t, 50, cnt.CountIP("10.0.0.0", "10.0.0.50"))
}

func Test1(t *testing.T) {
	var cnt IPCounter = countip.New()
	assert.Equal(t, 256, cnt.CountIP("10.0.0.0", "10.0.1.0"))
}

func Test2(t *testing.T) {
	var cnt IPCounter = countip.New()
	assert.Equal(t, 246, cnt.CountIP("20.0.0.10", "20.0.1.0"))
}
