package katas

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums    []int
		target  int
		indices []int
	}{
		{
			[]int{},
			0,
			nil,
		},
		{
			[]int{1},
			0,
			nil,
		},
		{
			[]int{3, 3, 3},
			5,
			nil,
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 2, 3},
			6,
			[]int{0, 2},
		},
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{

			[]int{-1, -2, -3, -4, -5},
			-8,
			[]int{2, 4},
		},
	}

	for _, test := range tests {
		indices := twoSum(test.nums, test.target)

		if !reflect.DeepEqual(indices, test.indices) {
			t.Errorf("for nums %v and target = %d -> indices = %v; want %v", test.nums, test.target, indices, test.indices)
		}
	}
}
