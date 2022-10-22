package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers445(l1 *ListNode, l2 *ListNode) *ListNode {

	reverse1 := reverseList(l1)
	reverse2 := reverseList(l2)

	dummyHead := &ListNode{}
	head := dummyHead
	carry := 0
	for reverse1 != nil || reverse2 != nil || carry > 0 {
		val := carry
		if reverse1 != nil {
			val = reverse1.Val + val
			reverse1 = reverse1.Next
		}
		if reverse2 != nil {
			val = reverse2.Val + val
			reverse2 = reverse2.Next
		}
		carry = val / 10
		head.Next = &ListNode{Val: val % 10}
		head = head.Next
	}
	return reverseList(dummyHead.Next)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}
