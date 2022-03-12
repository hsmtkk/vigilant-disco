package enoughenough_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/enoughenough"
	"github.com/stretchr/testify/assert"
)

/*
   assert_eq!(delete_nth(&[20,37,20,21], 1), vec![20,37,21]);
   assert_eq!(delete_nth(&[1,1,3,3,7,2,2,2,2], 3), vec![1, 1, 3, 3, 7, 2, 2, 2]);
*/

func Test0(t *testing.T) {
	want := []uint{20, 37, 21}
	input := []uint{20, 37, 20, 21}
	got := enoughenough.DeleteNth(input, 1)
	assert.Equal(t, want, got)
}

func Test1(t *testing.T) {
	want := []uint{1, 1, 3, 3, 7, 2, 2, 2}
	input := []uint{1, 1, 3, 3, 7, 2, 2, 2, 2}
	got := enoughenough.DeleteNth(input, 3)
	assert.Equal(t, want, got)
}
