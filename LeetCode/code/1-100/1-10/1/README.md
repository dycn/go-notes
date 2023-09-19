# [Two Sum](https://leetcode.com/problems/two-sum)

Given an array of integers, return **indices** of the two numbers such that they add up to a specific target.

You may assume that each input would have ***exactly*** one solution, and you may not use the *same*element twice.

**Example:**

```text
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```



[Code](p1.go)

[Test](p1_test.go)

## 思路
最优解法时间复杂度是O(n) 使用 map 将求解值与索引记录下来 循环遍历到需要这个求解值时，便可以获得对应的索引，避免两次for循环

HashTable
简单