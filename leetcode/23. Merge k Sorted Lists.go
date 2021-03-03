package main

import "algorithm/tool"

func init() {
	tool.Run(mergeKLists, []*ListNode{
		{1, &ListNode{3, &ListNode{Val: 5}}},
		{2, &ListNode{2, &ListNode{Val: 4}}},
	})
}

// 假设每个链表上有n个元素
// 如果每次循环仅合并一个，则比较的时间复杂度为O(nk^2)
// 如果两两合并则比较的时间复杂度为O(nklgk)
func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	for l > 1 {
		for i := 0; i < l; i += 2 {
			if i+1 < l {
				lists[i/2] = mergeTwo(lists[i], lists[i+1])
			} else {
				lists[i/2] = lists[i]
			}
		}
		l = (l + 1) / 2
		lists = lists[:l]
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
