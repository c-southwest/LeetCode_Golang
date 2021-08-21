package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	test := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	//test2:=&ListNode{Val: 1, Next: nil}
	r := removeNthFromEnd(test, 2)
	//r:=removeNthFromEnd(test2, 1)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var a, b, c *ListNode
	b = head
	c = b
	for i := 1; i < n; i++ {
		c = c.Next
	}
	// b,c 都就位了，开始前进
	for {
		if c.Next != nil {
			c = c.Next
			a = b // b移动之前需要用a来记录
			b = b.Next
		} else {
			//c到末尾了，此时b所指向的节点就应该被去掉
			if a != nil {
				a.Next = b.Next
				return head
			}
			//a和b都没移动过，所以返回的是 head.Next
			return b.Next
		}
	}
	return nil
}
