package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	assert.Equal(t, "AbC DeF", toWeirdCase("abc def"))
}

func Test1(t *testing.T) {
	assert.Equal(t, "AbC", toWeirdCase("ABC"))
}

func Test2(t *testing.T) {
	assert.Equal(t, "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd", toWeirdCase("This is a test Looks like you passed"))
}
