/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */
package main

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	l := len(nums) + 1
	for i, j := 0, 0; j < len(nums) || sum >= target; {
		if sum < target {
			sum += nums[j]
			j++
		} else {
			if j-i < l {
				l = j - i
			}
			sum -= nums[i]
			i++
		}
	}
	if l == len(nums)+1 {
		return 0
	} else {
		return l
	}
}

// @lc code=end
