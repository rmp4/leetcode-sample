package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question19 struct {
	para19
	ans19
}

type para19 struct {
	one []int
	n   int
}

type ans19 struct {
	one []int
}

func Test_Problem19(t *testing.T) {

	qs := []question19{

		{
			para19{[]int{1}, 3},
			ans19{[]int{1}},
		},

		{
			para19{[]int{1, 2}, 2},
			ans19{[]int{2}},
		},

		{
			para19{[]int{1}, 1},
			ans19{[]int{}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 10},
			ans19{[]int{1, 2, 3, 4, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 1},
			ans19{[]int{1, 2, 3, 4}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 2},
			ans19{[]int{1, 2, 3, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 3},
			ans19{[]int{1, 2, 4, 5}},
		},
		{
			para19{[]int{1, 2, 3, 4, 5}, 4},
			ans19{[]int{1, 3, 4, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 5},
			ans19{[]int{2, 3, 4, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 19------------------------\n")

	for _, q := range qs {
		_, p := q.ans19, q.para19
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(removeNthFromEnd(structures.Ints2List(p.one), p.n)))
	}
	fmt.Printf("\n\n\n")
}
