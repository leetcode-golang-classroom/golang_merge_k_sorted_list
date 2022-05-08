package sol

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKListsV1(lists []*ListNode) *ListNode {
	var head *ListNode
	Len := len(lists)
	if Len == 0 {
		return nil
	}
	if Len == 1 && lists[0] == nil {
		return nil
	}
	interval := 1
	for interval < Len {
		for idx := 0; idx < Len-1; idx += 2 * interval {
			if idx+interval < Len {
				lists[idx] = merge2List(lists[idx], lists[idx+interval])
			}
		}
		interval *= 2
	}
	head = lists[0]
	return head
}

func merge2List(l1, l2 *ListNode) *ListNode {
	var head, cur *ListNode
	for l1 != nil && l2 != nil {
		if head == nil {
			if l1.Val <= l2.Val {
				head = l1
				l1 = l1.Next
			} else {
				head = l2
				l2 = l2.Next
			}
			cur = head
		} else {
			if l1.Val <= l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		}
	}
	if head != nil {
		if l1 != nil {
			cur.Next = l1
		} else {
			cur.Next = l2
		}
	} else if l1 != nil {
		return l1
	} else {
		return l2
	}
	return head
}
