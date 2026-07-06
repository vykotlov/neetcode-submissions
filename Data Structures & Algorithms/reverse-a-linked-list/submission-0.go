func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	current := head.Next

	for current != nil {
		next := current.Next

		current.Next = prev

		prev = current
		current = next
	}

	head.Next = nil
	
	return prev
}
