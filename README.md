# golang_merge_k_sorted_list

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

*Merge all the linked-lists into one sorted linked-list and return it.*

## Examples

**Example 1:**

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

```

**Example 2:**

```
Input: lists = []
Output: []

```

**Example 3:**

```
Input: lists = [[]]
Output: []

```

**Constraints:**

- `k == lists.length`
- `0 <= k <= 104`
- `0 <= lists[i].length <= 500`
- `104 <= lists[i][j] <= 104`
- `lists[i]` is sorted in **ascending order**.
- The sum of `lists[i].length` will not exceed `104`.

## 解析

給定 k 個排序過由小排到大的 linked List 

要求要把 k 個排序過的 linked List 合成一個排序過的 linked List

最直覺的作法是把每次把 k 個 list 目前 head 找出最小的

逐步放入新的 linked List

這樣的每次要比k 次

而假設總共 n 個 時間複雜度為 O(k*N)

如果要把上述的時間複雜度做優化

可以透過 priority Queue 來讓 每次搜尋次數由 k 降低 logK

所以時間複雜度為 O(N*logK)

空間複雜度為 O(k) 來存放要比較的值

如果希望能把空間複雜度再降低到 O(1)

可以透過 Divide and Conquer的方式

要 merge k 個 linked List 先把 兩兩一對的 linked List 各自 merge

這樣 經過一次 merge 後只剩下 k/2 個

持續做下去可知一共需要 merge logK 次 而總供需要比較 N個

所以這樣的時間複雜度是 O(NlogK)

而空間複雜度是 O(1) ，不需要額外的空間來儲存比較的值

## 程式碼- 透過 prioryQueue

```go
import "container/heap"

type NodeHeap []*ListNode
func (h NodeHeap) Len() int {
    return len(h)
}
func (h NodeHeap) Less(i, j int) bool {
    return h[i].Val <= h[j].Val
}
func (h NodeHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *NodeHeap) Pop() interface{} {
    old := *h
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
    // loop for find current smallest
    var head *ListNode
    var currentNode *ListNode
    q := &NodeHeap{}
    for _, node := range lists {
        if node != nil {
            heap.Push(q, node)
        }
    }
    for q.Len() > 0 {
        n := heap.Pop(q).(*ListNode)
        
        if n.Next != nil {
            heap.Push(q, n.Next)
        }
        
        if (head == nil) {
            head = n
            currentNode = n
        } else {
            currentNode.Next = n
            currentNode = currentNode.Next
        }
    }
    return head
}
```

## 程式碼- Divide and Conquer

```go
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
```

## 困難點

1. 思考出降低比較次數的方式不是很直觀
2. Divide and Conquer 方式雖然不算難想，然而要看出其時間複雜度的優化需要良好的分析

## Solve Point

- [x]  Understand what problem would like to solve
- [x]  Analysis Complexity