# [15.3Sum](https://leetcode.com/problems/3sum/)

## Question
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
Note:
The solution set must not contain duplicate triplets.

Example:
```go
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 大意
給一個array 找到三個數之合等於0的所有組合
## 解題思路
先將資料排序 由小到大，利用兩個索引值 來做比較 而在這邊簡單判斷前後是否相同 如果相同就不處理j，k 两個索引值開始一前一後靠近。j為i 的下一個數字，k 為陣列的最後一個數字，由於經過排序，所以 k 的數字最大。j 往後移動，k 往前移動，逐漸找出最接近Target的值。

所使用的時間複雜度為O(n^2) 空間複雜度為Q(1)