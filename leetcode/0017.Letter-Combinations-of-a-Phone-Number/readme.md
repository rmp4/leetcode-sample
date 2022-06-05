# [17. Letter Combinations of a Phone Number]([link](https://leetcode.com/problems/letter-combinations-of-a-phone-number/))

## Question:
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

![alt text](200px-Telephone-keypad2.svg.png )

Example 1:
```go
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```
Example 2:
```go
Input: digits = ""
Output: []
```
Example 3:
```go
Input: digits = "2"
Output: ["a","b","c"]
```
Constraints:
```
0 <= digits.length <= 4
digits[i] is a digit in the range ['2', '9'].
```

## 大意
根據圖片所顯示的符號 輸入一字串 將所有輸入的字串對應的排列組合都顯示出來
## 解題思路
以DFS遞迴收尋即可