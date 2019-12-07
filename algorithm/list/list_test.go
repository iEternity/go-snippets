package list

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
	}
	answer := [][]int{
		{1},
		{2, 1},
		{3, 2, 1},
		{4, 3, 2, 1},
	}

	for i, nums := range cases {
		h := NewList(nums)
		h = Reverse(h)
		h2 := NewList(answer[i])
		if !Equal(h, h2) {
			t.Error(nums, answer[i])
		}
	}
}

func TestAdd(t *testing.T) {
	cases1 := [][]int{
		{9, 9, 9},
		{9, 9, 9},
		{1, 2, 3, 4},
	}
	cases2 := [][]int{
		{9, 9, 9},
		{9, 9, 9, 9},
		{1, 2, 0, 0},
	}
	answer := [][]int{
		{1, 9, 9, 8},
		{1, 0, 9, 9, 8},
		{2, 4, 3, 4},
	}

	for i, nums := range cases1 {
		h1, h2, h3 := NewList(nums), NewList(cases2[i]), NewList(answer[i])
		h := Add(h1, h2)
		if !Equal(h, h3) {
			t.Error(nums, cases2[i])
		}
	}
}
