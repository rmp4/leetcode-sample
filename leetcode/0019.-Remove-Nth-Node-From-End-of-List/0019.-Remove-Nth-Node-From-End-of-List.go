package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

type ListNode = structures.ListNode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{Next: head}

	previous, now, origin := res, head, head
	for origin != nil {
		if n <= 0 {
			previous = now
			now = now.Next
		}
		n--
		origin = origin.Next
	}
	previous.Next = now.Next
	return previous.Next
}
