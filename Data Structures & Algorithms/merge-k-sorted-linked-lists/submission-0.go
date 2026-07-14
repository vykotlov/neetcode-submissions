func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{}

	for _, secondList := range lists {
		mainList := dummy

		for secondList != nil && mainList.Next != nil {
			if mainList.Next.Val < secondList.Val {
				mainList = mainList.Next
				continue
			}

			mainListTmp := mainList.Next
			secondListTmp := secondList.Next

			mainList.Next = secondList
			mainList.Next.Next = mainListTmp

			secondList = secondListTmp
		}

		if mainList.Next == nil {
			mainList.Next = secondList
		}
	}

	return dummy.Next
}
