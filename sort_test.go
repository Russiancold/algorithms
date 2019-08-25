package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct{
		in []int
		sorted []int
	} {
		{[]int{5,4,1,8,7,2,6,3}, []int{1,2,3,4,5,6,7,8}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{3,2}, []int{2,3}},
		{[]int{2,1,3}, []int{1,2,3}},
	}

	for _, tc := range testCases {
		assert.Equal(tc.sorted, MergeSort(tc.in))
	}
}
