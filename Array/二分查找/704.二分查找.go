/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := (l + r) / 2
	for l <= r {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
			mid = (l + r) / 2
		} else {
			l = mid + 1
			mid = (l + r) / 2
		}
	}
	return -1
}

// @lc code=end

