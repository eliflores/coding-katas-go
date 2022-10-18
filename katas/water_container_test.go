package katas

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		area   int
	}{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 8, 6, 2, 5, 7, 8, 3, 4},
			40,
		},
		{
			[]int{1, 8, 6, 2, 5, 7, 4, 3, 4},
			28,
		},
		{
			[]int{2, 3, 4, 5, 18, 17, 6},
			17,
		},
		{
			[]int{1, 2, 1},
			2,
		},
		{
			[]int{1, 1},
			1,
		},
	}

	for _, test := range tests {
		area := maxArea(test.height)

		if area != test.area {
			t.Errorf("area height = %d; want %d", area, test.area)
		}
	}
}
