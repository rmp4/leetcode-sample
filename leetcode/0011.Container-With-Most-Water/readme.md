#08 . [11. Container With Most Water]([(https://leetcode.com/problems/string-to-integer-atoi/](https://leetcode.com/problems/container-with-most-water/)))

## Question:
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
```go
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
```
Example 2:
```go
Input: height = [1,1]
Output: 1
```


Constraints:
```
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
```

## 大意
會輸入一組正整數的陣列 需要找出其最大面積 其中面積是根據陣列的最大值以及其x軸所組成的
## 解題思路

先從 x 軸下手 以倆個指針分別為0(start)以及陣列最大值(end) 開始往內縮，而指針更改條件為找到最大跟第二大的值，透過這樣來將以第二高的值算出期面積
