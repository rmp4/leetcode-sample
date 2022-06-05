package leetcode

import (
	"math"
	"sort"
)

// O(n^2)
func threeSumClosest(nums []int, target int) int {
	res := 0
	length := len(nums)
	diff := math.MaxInt32
	if length > 2 {
		sort.Ints(nums)
		for i := 0; i < length-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, length-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(number int) int {
	if number > 0 {
		return number
	}
	return -number
}
