package unluckydays_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/unluckydays"
	"github.com/stretchr/testify/assert"
)

func Test1986(t *testing.T) {
	got := unluckydays.New().Calculate(1986)
	const want uint = 1
	assert.Equal(t, want, got)
}

func Test2015(t *testing.T) {
	got := unluckydays.New().Calculate(2015)
	const want uint = 3
	assert.Equal(t, want, got)
}
