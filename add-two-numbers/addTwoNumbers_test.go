package addtwonumbers

import (
	"reflect"
	"testing"
)

func array2ListNode(values []int) *ListNode {
	tmp := &ListNode{Val: values[0]}
	cur := tmp

	for _, val := range values {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return tmp.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				l1: array2ListNode([]int{2, 4, 3}),
				l2: array2ListNode([]int{5, 6, 4}),
			},
			want: array2ListNode([]int{7, 0, 8}),
		},
		{
			name: "Case 2",
			args: args{
				l1: array2ListNode([]int{1, 8}),
				l2: array2ListNode([]int{0}),
			},
			want: array2ListNode([]int{1, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
