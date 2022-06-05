package leetcode

import (
	"fmt"
	"testing"
)

type question88 struct {
	para88
	ans88
}
type para88 struct {
	one []int
	m   int
	two []int
	n   int
}

type ans88 struct {
	one []int
}

func Test_problem88(t *testing.T) {
	qs := []question88{
		{
			para88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans88{[]int{1, 2, 2, 3, 5, 6}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")
	for _, q := range qs {
		_, p := q.ans88, q.para88
		fmt.Printf("【input】:%v,%v,%v,%v       ", p.one, p.m, p.two, p.n)
		merge(p.one, p.m, p.two, p.n)
		fmt.Printf("【output】:%v\n", p)
	}
	fmt.Printf("\n\n\n")
}
