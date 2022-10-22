package leetcode

import (
	"github.com/rmp4/leetcode-sample/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

func zigzagLevelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	// Initial variable
	size := 0
	i, j := 0, 0
	lay := []int{}
	tmp := []*TreeNode{}
	flag := false

	for len(q) > 0 {
		size = len(q)
		tmp = []*TreeNode{}
		lay = make([]int, size)
		j = size - 1
		for i = 0; i < size; i++ {
			root = q[0]
			q = q[1:]
			if !flag {
				lay[i] = root.Val
			} else {
				lay[j] = root.Val
				j--
			}
			if root.Left != nil {
				tmp = append(tmp, root.Left)
			}
			if root.Right != nil {
				tmp = append(tmp, root.Right)
			}

		}
		res = append(res, lay)
		flag = !flag
		q = tmp
	}
	return res
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := NewQueue()
	queue.Enqueue(root)
	reverse := false
	for !queue.Empty() {
		count := queue.Count()
		resItems := make([]int, 0)
		for count > 0 {
			node := queue.Dequeue()
			resItems = append(resItems, node.Val)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
			if count == 1 {
				if reverse {
					reverseItems(resItems)
				}
				res = append(res, resItems)
				reverse = !reverse
			}
			count -= 1
		}
	}
	return res
}

func reverseItems(items []int) {
	i := 0
	j := len(items) - 1
	for i < j {
		items[i], items[j] = items[j], items[i]
		i += 1
		j -= 1
	}
}

type Queue struct {
	items []*TreeNode
}

func (self *Queue) Enqueue(node *TreeNode) {
	self.items = append(self.items, node)
}

func (self *Queue) Dequeue() *TreeNode {
	if len(self.items) == 0 {
		return nil
	}
	node := self.items[0]
	self.items = self.items[1:]
	return node
}

func (self *Queue) Count() int {
	return len(self.items)
}

func (self *Queue) Empty() bool {
	return len(self.items) == 0
}

func NewQueue() Queue {
	return Queue{
		items: make([]*TreeNode, 0),
	}
}
