package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dummy := &Node{-1, nil, nil}
	cur := dummy
	first := head

	for ; first != nil; {
		cur.Next = first
		cur = first
		tmp := &Node{first.Val, nil, nil}
		tmp.Next = first.Next
		cur.Next = tmp
		cur = cur.Next
		first = first.Next.Next
	}

	first = head

	for ; first != nil; {
		if first.Random != nil {
			first.Next.Random = first.Random.Next
		} else {
			first.Next.Random = nil
		}

		first = first.Next.Next
	}

	first = head
	cur = dummy
	for ; first != nil; {
		cur.Next = first.Next
		first.Next = first.Next.Next
		first = first.Next
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	n13 := &Node{13,nil,nil}
	head := &Node{7,n13,nil}
	n13.Random = head
	copyRandomList(head)
}
