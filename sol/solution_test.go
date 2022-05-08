package sol

import (
	"reflect"
	"testing"
)

func CreateLists(nums *[][]int) []*ListNode {
	arr := *nums
	lists := make([]*ListNode, len(arr))
	for idx, nodes := range arr {
		lists[idx] = CreateList(&nodes)
	}
	return lists
}
func CreateList(nums *[]int) *ListNode {
	var head, cur *ListNode
	arr := *nums
	for idx, val := range arr {
		if idx == 0 {
			head = &ListNode{Val: val}
			cur = head
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
	}
	return head
}
func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "lists = [[1,4,5],[1,3,4],[2,6]]",
			args: args{lists: CreateLists(&[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})},
			want: CreateList(&[]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "lists = []",
			args: args{lists: CreateLists(&[][]int{})},
			want: CreateList(&[]int{}),
		},
		{
			name: "lists = [[]]",
			args: args{lists: CreateLists(&[][]int{{}})},
			want: CreateList(&[]int{}),
		},
		{
			name: "lists = [[],[]]",
			args: args{lists: CreateLists(&[][]int{{}, {}})},
			want: CreateList(&[]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
