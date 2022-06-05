# [19. Remove Nth Node From End of List]([link](https://leetcode.com/problems/remove-nth-node-from-end-of-list/))

## Question:
Given the head of a linked list, remove the nth node from the end of the list and return its head.

 ![alt text](remove_ex1.jpg)

Example 1:

```go
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```
Example 2:
```go
Input: head = [1,2], n = 1
Output: [1]
```

Constraints:
```
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
```

## 大意
移除Linklist從後面數來的n個node
## 解題思路
先複製一個原本的Link list 並且重頭走過一次，每往下一個節點時將n-1當n == 0時表示要移除掉該節點並解將之後的節點重新連接到下一個位置
當原本的Linklist走完後也處理完所有的節點