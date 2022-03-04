package diffvolcube_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/diffvolcube"
	"github.com/stretchr/testify/assert"
)

func Test0(t *testing.T) {
	const want uint = 14
	got := diffvolcube.CalculateDifference([]uint{3, 2, 5}, []uint{1, 4, 4})
	assert.Equal(t, want, got)
	/*
			        assert_eq!(find_difference(&[3, 2, 5], &[1, 4, 4]), 14);
		        assert_eq!(find_difference(&[9, 7, 2], &[5, 2, 2]), 106);

	*/
}

func Test1(t *testing.T) {
	const want uint = 106
	got := diffvolcube.CalculateDifference([]uint{9, 7, 2}, []uint{5, 2, 2})
	assert.Equal(t, want, got)
}
