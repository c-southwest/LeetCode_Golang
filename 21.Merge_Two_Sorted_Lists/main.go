package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, p *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	p = head
	for {
		if l1 == nil {
			p.Next = l2
			return head
		} else if l2 == nil {
			p.Next = l1
			return head
		}

		if l1.Val < l2.Val {
			p.Next = l1
			p = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			p = l2
			l2 = l2.Next
		}
	}
}
