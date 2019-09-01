package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func listNode2Array(l *ListNode) []int {
	arr := make([]int, 0)
	for l != nil {
		arr = append(arr, l.Val)
		l = l.Next
	}
	return arr
}

type input struct {
	l1 *ListNode
	l2 *ListNode
}

// Case test definition
type Case struct {
	parameters input
	expected   []int
}

func TestAddTwoNumbers(t *testing.T) {
	cases := []Case{
		Case{
			parameters: input{
				l1: array2ListNode([]int{2, 4, 3}),
				l2: array2ListNode([]int{5, 6, 4}),
			},
			expected: []int{7, 0, 8},
		},
		Case{
			parameters: input{
				l1: array2ListNode([]int{1, 8}),
				l2: array2ListNode([]int{0}),
			},
			expected: []int{1, 8},
		},
	}

	for _, c := range cases {
		result := addTwoNumbers(c.parameters.l1, c.parameters.l2)
		assert.Equal(t, c.expected, listNode2Array(result))
	}
}
