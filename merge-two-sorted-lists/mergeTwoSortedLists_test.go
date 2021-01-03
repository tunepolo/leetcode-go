package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func array2ListNode(values []int) *ListNode {
	tmp := &ListNode{Val: 0}
	cur := tmp

	for _, val := range values {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return tmp.Next
}

func listNode2Array(l *ListNode) []int {
	arr := make([]int, 0)
	for l != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return arr
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				l1: array2ListNode([]int{1, 2, 4}),
				l2: array2ListNode([]int{1, 3, 4}),
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "Case 2",
			args: args{
				l1: array2ListNode([]int{}),
				l2: array2ListNode([]int{}),
			},
			want: []int{},
		},
		{
			name: "Case 3",
			args: args{
				l1: array2ListNode([]int{}),
				l2: array2ListNode([]int{0}),
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listNode2Array(mergeTwoLists(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
