package sol

import (
	"reflect"
	"testing"
)

func Test_mergeKListsV1(t *testing.T) {
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
		{
			name: "lists = [[],[1]]",
			args: args{lists: CreateLists(&[][]int{{}, {1}})},
			want: CreateList(&[]int{1}),
		},
		{
			name: "lists = [[],[],[],[],[],[],[],[],[],[],[]]",
			args: args{lists: CreateLists(&[][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}})},
			want: CreateList(&[]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListsV1(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKListsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
