# [680. Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/)


Given a string s, return true if the s can be palindrome after deleting at most one character from it.

 

Example 1:

Input: s = "aba"
Output: true
Example 2:

Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.
Example 3:

Input: s = "abc"
Output: false


Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.



[Code](p680.go)

[Test](p680_test.go)

## 思路
1. 穷举所有的字符串并每个判断是否是回文