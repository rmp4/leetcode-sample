package leetcode

import (
	"fmt"
	"testing"
)

type question445 struct {
	para445
	ans445
}

// para 是参数
// one 代表第一个参数
type para445 struct {
	one     []int
	another []int
}

type ans445 struct {
	one []int
}

func Test_Problem445(t *testing.T) {

	qs := []question445{

		{
			para445{[]int{}, []int{}},
			ans445{[]int{}},
		},

		{
			para445{[]int{1}, []int{1}},
			ans445{[]int{2}},
		},

		{
			para445{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans445{[]int{2, 4, 6, 8}},
		},

		{
			para445{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans445{[]int{2, 4, 6, 9, 0}},
		},

		{
			para445{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		{
			para445{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		{
			para445{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans445{[]int{8, 0, 7}},
		},

		{
			para445{[]int{1, 8, 3}, []int{7, 1}},
			ans445{[]int{2, 5, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 445------------------------\n")

	for _, q := range qs {
		_, p := q.ans445, q.para445
		fmt.Printf("【input】:%v       【output】:%v\n", p, List2Ints(addTwoNumbers445(Ints2List(p.one), Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
