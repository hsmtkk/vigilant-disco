package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/comsqstr/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	want := "bNkTB\nhTrWO\nRTFVi\nCnnIj"
	got := kata.Compose("byGt\nhTts\nRTFF\nCnnI", "jIRl\nViBu\nrWOb\nNkTB")
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := "HgYPW\nTGGbM\nIPhqt\nuUMDH"
	got := kata.Compose("HXxA\nTGBf\nIPhg\nuUMD", "Hcbj\nqteH\nGbMJ\ngYPW")
	assert.Equal(t, want, got)
}
