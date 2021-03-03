package main

func mergeKLists(lists []*ListNode) *ListNode {
	for len(lists) > 1 {
		l := len(lists)
		for i := 0; i < l; i += 2 {
			if i+1 < l {
				lists[i/2] = mergeTwo(lists[i], lists[i+1])
			} else {
				lists[i/2] = lists[i]
			}
		}
		lists = lists[:(l+1)/2]
	}
	return lists[0]
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	h := &ListNode{}
	cur := h
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return h.Next
}
