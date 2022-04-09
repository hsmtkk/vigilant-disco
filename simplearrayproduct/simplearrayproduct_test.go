package simplearrayproduct_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/simplearrayproduct"
	"github.com/stretchr/testify/assert"
)

/*
   assert_eq!(
       solve(&[vec![-11, -6], vec![-20, -20], vec![18, -4], vec![-20, 1]]),
       17600
   );
   assert_eq!(solve(&[vec![-3, -4], vec![1, 2, -3]]), 12);

*/

func Test0(t *testing.T) {
	assert.Equal(t, 8, simplearrayproduct.Solve([]int{1, 2}, []int{3, 4}))
}

func Test1(t *testing.T) {
	assert.Equal(t, 45, simplearrayproduct.Solve([]int{10, -15}, []int{-1, -3}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 12, simplearrayproduct.Solve([]int{-1, 2, -3, 4}, []int{1, -2, 3, -4}))
}

func Test3(t *testing.T) {
	assert.Equal(t, 3584, simplearrayproduct.Solve([]int{14, 2}, []int{0, -16}, []int{-12, -16}))
}
