package LinkedList


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	// 虚拟节点头，防止出现链表全部被删除的情况
	dummy := &ListNode{-1, nil}
	dummy.Next = head
	cur := dummy
	//限移动一部分
	for i := 0; i < n; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	//pre移动到要删除的节点前方
	pre := dummy
	for ; cur.Next != nil; {
		cur = cur.Next
		pre = pre.Next
	}
	//删除倒数第n个节点
	pre.Next = pre.Next.Next
	return dummy.Next
}
