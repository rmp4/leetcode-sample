# 16. [3Sum Closest](https://leetcode.com/problems/3sum-closest/)

## Question:
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.

Example 1:
```go
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
```
Example 2:
```go
Input: nums = [0,0,0], target = 1
Output: 0
```

Constraints:
```
3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
```

## 大意
給一個array 找到三個數之合 距離Target最接近
## 解題思路
先將資料排序 由小到大，利用兩個索引值 來做比較 而在這邊簡單判斷前後是否相同 如果相同就不處理j，k 两個索引值開始一前一後靠近。j為i 的下一個數字，k 為陣列的最後一個數字，由於經過排序，所以 k 的數字最大。j 往後移動，k 往前移動，逐漸找出最接近Target的值。