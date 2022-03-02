package muldigsum_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/muldigsum"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	const want uint = 72
	got := muldigsum.MulDigSum(12)
	assert.Equal(t, want, got)
}

func TestMultiples(t *testing.T) {
	want := []uint{12, 24, 36, 48, 60, 72, 84, 96}
	got := muldigsum.Multiples(12)
	assert.Equal(t, want, got)
}

func TestDigitSum(t *testing.T) {
	const want uint = 3
	got := muldigsum.DigitSum(12)
	assert.Equal(t, want, got)
}
