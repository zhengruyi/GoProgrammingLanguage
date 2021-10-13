package Other

type ListNode struct {
	Val  int
	Next *ListNode
}

//将下一节点的值复制到当前节点,然后将下一节点删除就可以了
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
