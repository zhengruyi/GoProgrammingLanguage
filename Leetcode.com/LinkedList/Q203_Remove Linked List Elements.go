package LinkedList

// 重点在于虚拟节点的处理
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{-1,head}
	cur := dummy
	for ;cur != nil && cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return dummy.Next
}