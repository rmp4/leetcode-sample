package leetcode

import (
	"fmt"
	"testing"
)

type question16 struct {
	para16
	ans16
}

// para is parameter
// one is first parameter
type para16 struct {
	a      []int
	target int
}

// ans is answer
// one is first answer
type ans16 struct {
	one int
}

func Test_Problem16(t *testing.T) {
	qs := []question16{
		{
			para16{[]int{-1, 0, 1, 1, 55}, 3},
			ans16{2},
		},
		{
			para16{[]int{0, 0, 0}, 1},
			ans16{0},
		},
		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},
		{
			para16{[]int{1, 1, -1}, 0},
			ans16{1},
		},
		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 16------------------------\n")
	for _, q := range qs {
		_, p := q.ans16, q.para16
		fmt.Printf("【input】:%v       【output】:%v\n", p, threeSumClosest(p.a, p.target))
	}
	fmt.Printf("\n\n\n")
}
