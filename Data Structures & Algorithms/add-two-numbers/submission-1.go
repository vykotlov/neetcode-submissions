func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultDummy := result
	shift := false

	for l1 != nil || l2 != nil {
		left, right := 0, 0

		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}

		sum := left + right

		if shift {
			sum++
			shift = false
		}

		if sum >= 10 {
			sum = sum - 10
			shift = true
		}

		result.Next = &ListNode{Val: sum}
		result = result.Next
	}

	if shift {
		result.Next = &ListNode{Val: 1}
	}

	return resultDummy.Next
}
