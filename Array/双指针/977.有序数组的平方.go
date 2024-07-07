/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
package main

// @lc code=start
func sortedSquares(nums []int) []int {
	l, r, index := 0, len(nums)-1, len(nums)-1
	res := make([]int, len(nums))
	for l <= r {
		lSquare, rSquare := nums[l]*nums[l], nums[r]*nums[r]
		if lSquare <= rSquare {
			res[index] = rSquare
			r--
		} else {
			res[index] = lSquare
			l++
		}
		index--
	}
	return res
}

// @lc code=end
