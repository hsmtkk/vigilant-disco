package kata_test

import(
	"testing"
	"github.com/hsmtkk/vigilant-disco/oddnum/kata"
	"github.com/stretchr/testify/assert"
)

func Test15(t *testing.T){
	assert.Equal(t, 7, kata.OddCount(15))	
}

func Test15023(t *testing.T){
	assert.Equal(t, 7511, kata.OddCount(15023))	
}
