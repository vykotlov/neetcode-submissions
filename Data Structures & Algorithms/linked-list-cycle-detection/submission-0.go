func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	
	pointer1 := head
	pointer2 := head
	
	for {
		pointer1 = pointer1.Next
		if pointer1 == nil {
			break
		}
		
		pointer2 = pointer2.Next
		if pointer2 == nil {
			break
		}
		
		pointer2 = pointer2.Next
		if pointer2 == nil {
			break
		}
		
		if pointer1 == pointer2 {
			return true
		}
	}
	
	return false
}
