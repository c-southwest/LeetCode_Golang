package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
//	result := ListNode{}
//	var num1, index1 int64 = 0, 1
//
//	for l1 != nil {
//		num1 += l1.Val * index1
//		l1 = l1.Next
//		index1 *= 10
//	}
//
//	num2, index2 := 0, 1
//	for l2 != nil {
//		num2 += l2.Val * index2
//		l2 = l2.Next
//		index2 *= 10
//	}
//
//	sum := num1 + num2
//	pointer := &result
//	for {
//		if sum < 10 {
//			pointer.Val = sum
//			pointer.Next = nil
//			return &result
//		}
//		mod := sum % 10
//		sum = sum / 10
//		pointer.Val = mod
//		pointer.Next = &ListNode{Val: 0, Next: nil}
//		pointer = pointer.Next
//	}
//
//}

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 2, Next: nil}}
	l2 := &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: nil}}
	r := addTwoNumbers(l1, l2)
	println(r)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{}
	var nums1, nums2, sums []int
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next
	}
	for {
		if len(nums1) > len(nums2) {
			for k, v := range nums1 {
				if k < len(nums2) {
					sums = append(sums, v+nums2[k])
				} else {
					sums = append(sums, v)
				}
			}
			break
		} else {
			for k, v := range nums2 {
				if k < len(nums1) {
					sums = append(sums, v+nums1[k])
				} else {
					sums = append(sums, v)
				}
			}
			break
		}
	}
	pointer := &result
	flag := 0
	for i, v := range sums {
		if flag == 1 {
			v += 1
			flag = 0
		}
		if v >= 10 {
			flag = 1
			v -= 10
		}
		pointer.Val = v
		if i == len(sums)-1 {
			if flag == 1 {
				pointer.Next = &ListNode{Next: nil}
				pointer = pointer.Next
				pointer.Val = 1
				break
			} else {
				break
			}

		}
		pointer.Next = &ListNode{Next: nil}
		pointer = pointer.Next
	}

	return &result

}
