/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	for curA != nil {
		lenA++
		curA = curA.Next
	}
	for curB != nil {
		lenB++
		curB = curB.Next
	}
	var step int
	var slow, fast *ListNode
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast

}

// @lc code=end
