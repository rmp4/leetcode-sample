package leetcode

import (
	"fmt"
	"testing"

	"github.com/rmp4/leetcode-sample/structures"
)

type question103 struct {
	para103
	ans103
}

// para is parameter
// one is first parameter
type para103 struct {
	one []int
}

// ans is answer
// one is first answer
type ans103 struct {
	one [][]int
}

func Test_Problem103(t *testing.T) {
	qs := []question103{
		{
			para103{[]int{}},
			ans103{[][]int{}},
		},

		{
			para103{[]int{1}},
			ans103{[][]int{{1}}},
		},

		{
			para103{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans103{[][]int{{3}, {9, 20}, {15, 7}}},
		},

		{
			para103{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, 5}},
			ans103{[][]int{{1}, {3, 2}, {4, 5}}},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 103------------------------\n")

	for _, q := range qs {
		_, p := q.ans103, q.para103
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", zigzagLevelOrder(root))
	}
	fmt.Printf("\n\n\n")
}
