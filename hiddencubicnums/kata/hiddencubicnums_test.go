package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/hiddencubicnums/kata"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "0 0 Lucky", kata.IsSumOfCubes("0 9026315 -827&()"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "Unlucky", kata.IsSumOfCubes("Once upon a midnight dreary, while100 I pondered, 9026315weak and weary -827&()"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "0 0 Lucky", kata.IsSumOfCubes("Once 1000upon a midnight 110dreary, while100 I pondered, 9026315weak and weary -827&()"))
}
