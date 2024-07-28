package main

import (
	"fmt"

	linkedlist "github.com/P-Jstone/Algorithm/Linkedlist"
)

func main() {
	ll := &linkedlist.ListNode{
		Val:  1,
		Next: nil,
	}
	res := linkedlist.RemoveNthFromEnd(ll, 1)
	fmt.Println(res)
}
