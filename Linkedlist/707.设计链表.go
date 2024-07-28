/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */
package linkedlist

// @lc code=start
type MyLinkedList struct {
	dummyHead *Node
	size      int
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	head := &Node{
		-1,
		nil,
	}
	return MyLinkedList{
		dummyHead: head,
		size:      0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size {
		return -1
	}
	node := this.dummyHead.Next
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: this.dummyHead.Next,
	}
	this.dummyHead.Next = node
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	tailNode := &Node{
		Val:  val,
		Next: nil,
	}
	node := this.dummyHead
	for node.Next != nil {
		node = node.Next
	}
	node.Next = tailNode
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index > this.size {
		return
	}
	node := this.dummyHead
	for i := 0; i < index; i++ {
		node = node.Next
	}
	node.Next = &Node{
		Val:  val,
		Next: node.Next,
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	node := this.dummyHead
	for i := 0; i < index; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
