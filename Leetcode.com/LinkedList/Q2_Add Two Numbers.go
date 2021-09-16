package LinkedList

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	node := dummy
	carry := 0
	val := 0
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			val = (carry + l1.Val + l2.Val) % 10
			carry = (carry + l1.Val + l2.Val) / 10
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			val = (carry + l1.Val) % 10
			carry = (carry + l1.Val) / 10
			l1 = l1.Next
		} else if l2 != nil {
			val = (carry + l2.Val) % 10
			carry = (carry + l2.Val) / 10
			l2 = l2.Next
		}
		node.Next = &ListNode{val, nil}
		node = node.Next
	}
	if carry != 0 {
		node.Next = &ListNode{1, nil}
	}
	return dummy.Next
}
