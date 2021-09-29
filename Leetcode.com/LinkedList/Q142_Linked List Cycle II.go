package LinkedList

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	fast := head
	slow := head

	for; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 尤其记住第一次相遇后，快节点也变成每次一步，这样再次相遇的节点就是入口节点
	slow = head
	for; fast != slow; {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}