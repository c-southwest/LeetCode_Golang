package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil {
		tmp := p
		v := p.Val

		p = p.Next
		if p != nil {
			if v == p.Val {
				tmp.Next = p.Next
				p = tmp
				continue
			} else {
				continue
			}
		} else {
			return head
		}
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	v := 101
	if head != nil {
		v = head.Val
	} else {
		return head
	}
	pre := head
	next := pre.Next

	for next != nil {
		if next.Val == v {
			pre.Next = next.Next
			next = pre.Next
		} else {
			v = next.Val
			pre = next
			next = pre.Next
		}
	}

	return head
}
