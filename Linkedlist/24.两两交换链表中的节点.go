/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		pre = head
		head = next
	}
	return dummy.Next
}

// @lc code=end
