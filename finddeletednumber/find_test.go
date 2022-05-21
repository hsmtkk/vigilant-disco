package finddeletednumber_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/finddeletednumber"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	list          []int
	mixedList     []int
	deleted       bool
	deletedNumber int
}

func TestFind(t *testing.T) {
	testCases := []testCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 2, 4, 6, 7, 8, 1, 9, 5}, false, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 2, 4, 6, 7, 8, 1, 9}, true, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 2, 4, 6, 7, 8, 9, 5}, true, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 2, 4, 1, 7, 8, 9, 5}, true, 6},
	}
	for _, tc := range testCases {
		gotDeleted, gotDeletedNumber := finddeletednumber.Find(tc.list, tc.mixedList)
		assert.Equal(t, tc.deleted, gotDeleted)
		assert.Equal(t, tc.deletedNumber, gotDeletedNumber)
	}
}
