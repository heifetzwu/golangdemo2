package main

import (
	"fmt"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers_main() {
	fmt.Println("addTwoNumbers_main")
	l1_head := new(ListNode)
	l1_head.Val = 8

	l1 := new(ListNode)
	l1.Val = 2
	l1_head.Next = l1

	l2_head := new(ListNode)
	l2_head.Val = 8

	l2 := new(ListNode)
	l2.Val = 2
	l2_head.Next = l1

	lr := addTwoNumbers(l1_head, l2_head)
	for l := lr; true; {

		fmt.Println(l.Val)
		if l.Next == nil {
			time.Sleep(1 * time.Second)
			break
		}
		l = l.Next
	}

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode
	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		v := l1.Val + l2.Val + carry
		carry = v / 10
		v = v % 10

		if head == nil {
			head = &ListNode{Val: v, Next: nil}
			cur = head
		} else {
			cur.Next = &ListNode{v, nil}
			cur = cur.Next
		}

	}
	if l1 != nil || l2 != nil {
		if l2 != nil {
			l1 = l2
		}
		for ; l1 != nil; l1 = l1.Next {
			v := l1.Val + carry
			carry = v / 10
			v = v % 10

			if head == nil {
				head = &ListNode{v, nil}
				cur = head
			} else {
				cur.Next = &ListNode{v, nil}
				cur = cur.Next
			}
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return head
}
