package sol

import "container/heap"

type NodeHeap []*ListNode

func (h *NodeHeap) Len() int {
	return len(*h)
}

func (h *NodeHeap) Less(i, j int) bool {
	return (*h)[i].Val <= (*h)[j].Val
}

func (h *NodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	// pop most top
	value := old[len(old)-1]
	*h = old[:len(old)-1]
	return value
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var curNode *ListNode
	q := &NodeHeap{}
	// push all head node into q
	for _, node := range lists {
		if node != nil {
			heap.Push(q, node)
		}
	}
	for q.Len() > 0 {
		// O(1)
		n := heap.Pop(q).(*ListNode)
		if n.Next != nil {
			// O(logK)
			heap.Push(q, n.Next)
		}
		if head == nil {
			head = n
			curNode = head
		} else {
			curNode.Next = n
			curNode = curNode.Next
		}
	}
	return head
}
