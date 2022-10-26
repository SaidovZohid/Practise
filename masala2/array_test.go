package masala2

import (
	"testing"
)

func TestArray(t *testing.T) {
	TestCases := [][]int{
		{1, 2, 3, 1, 4, 5, 6, 7, 8, 0, -7, -1, -7, 23, 32344, -7},
		{3, 4, 9, 8, 3, 5, -1, -6, 234, -6, 12, -6},
		{3, 4, 9, 8, 32, 12, 65, 3, 90},
		{3, 4, 9, 8, 3, -1, 43, 12, 98, -1},
		{3, 4, 9, 8, 3, 0, 1, -13, 233, -13},
		{8, 3, 0, 1, 5, 7, 1, 9, 0, -1, -1, -4},
		{3, 4, 9, 8, 3},
		{3, 4, 9, 8, 3, 0, 1},
	}
	for _, testcase := range TestCases {
		res := FindMinimalElements(testcase)
		t.Log(res)
	}
}