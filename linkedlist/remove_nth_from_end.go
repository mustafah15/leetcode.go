package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// creating a dummy node to hold the head
	vhead := &ListNode{Next: head}
	// two pointers pointing to the dummy head
	l, r := vhead, vhead

	// moving one pointer with gap of n
	for n >= 0 {
		r = r.Next
		n--
	}

	// moving toward the end of the list
	// when we reach the end with r
	// we reach the nth node
	for r != nil {
		r = r.Next
		l = l.Next
	}

	// remove the nth node
	l.Next = l.Next.Next

	// return the start of the linked list
	return vhead.Next
}
