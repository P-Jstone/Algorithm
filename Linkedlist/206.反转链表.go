/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
package linkedlist

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = pre
		pre = temp
	}
	return pre
}

// @lc code=end
