# [6. ZigZag Conversio](https://leetcode.com/problems/zigzag-conversion/)

## Question:


The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string s, int numRows);
```


Example 1:

```go
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

Example 2:

```go
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     IInput: nums = [0,0,0], target = 1
Output: 0
```

Example 3:

```go
Input: s = "A", numRows = 1
Output: "A"
```

Constraints:

```
1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
```

## 大意

將一字串Ｓ輸入後根據所輸入的整數 numRows 從上往下且由左而右排列成Ｚ字型

比如輸入為 `"PINALSIGYAHRPI"` 且numRows 為３時排列如下：

```go
P   A   H   N
A P L S I I G
Y   I   R
```

之後將輸出字串由左讀到右 並且產生一個新的字串 如 `"PAHNAPLSIIGYIR"`

## 解題思路

利用兩個變數來儲存方向，當垂直輸出的行數到達規定的行數時需要從下開始往上到第一行
