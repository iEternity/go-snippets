package list

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func NewList(nums []int) *Node {
	root := &Node{Val: nums[0]}
	prev := root
	for i := 1; i < len(nums); i++ {
		prev.Next = &Node{Val: nums[i]}
		prev = prev.Next
	}
	return root
}

func PrintList(h *Node) {
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func Equal(h1, h2 *Node) bool {
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1, h2 = h1.Next, h2.Next
	}

	if h1 != nil {
		return false
	}
	if h2 != nil {
		return false
	}
	return true
}

func Reverse(h *Node) *Node {
	var prev *Node
	for h != nil {
		h.Next, h, prev = prev, h.Next, h
	}
	return prev
}

func Add(h1, h2 *Node) *Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	h1 = Reverse(h1)
	h2 = Reverse(h2)

	var prev, head *Node
	r := 0
	for h1 != nil && h2 != nil {
		n := h1.Val + h2.Val + r
		r = n / 10
		if prev == nil {
			head = &Node{Val: n % 10}
			prev = head
		} else {
			prev.Next = &Node{Val: n % 10}
			prev = prev.Next
		}

		h1, h2 = h1.Next, h2.Next
	}

	for h1 != nil {
		n := h1.Val + r
		r = n / 10
		prev.Next = &Node{Val: n % 10}
		prev, h1 = prev.Next, h1.Next
	}

	for h2 != nil {
		n := h2.Val + r
		r = n / 10
		prev.Next = &Node{Val: n % 10}
		prev, h2 = prev.Next, h2.Next
	}

	if r != 0 {
		prev.Next = &Node{Val: r}
	}

	return Reverse(head)
}
