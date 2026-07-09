func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var current *ListNode

	for list1 != nil || list2 != nil {
		currentValue := 0

		if list2 == nil || list1 != nil && list1.Val < list2.Val {
			currentValue = list1.Val
			list1 = list1.Next
		} else {
			currentValue = list2.Val
			list2 = list2.Next
		}
		
		next := &ListNode{
			Val: currentValue,
			Next: nil,
		}

		if current == nil {
			head = next
			current = next
		} else {
			current.Next = next
			current = next
		}
	}

	return head
}