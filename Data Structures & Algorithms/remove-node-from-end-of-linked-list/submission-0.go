func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	
	current := head
	count := 0

	for current != nil {
		current = current.Next
		count++
	}
	
	removeAt := count - n
	if removeAt < 0 {
		return head
	}
	
	if removeAt == 0 {
		return head.Next
	}

	current = head
	for i := 0; i < removeAt - 1; i++ {
		current = current.Next
	}
	
	current.Next = current.Next.Next
	return head
}