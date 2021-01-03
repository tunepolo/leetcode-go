package mergetwosortedlists

// https://leetcode.com/problems/merge-two-sorted-lists/

// ListNode is linked-list node that store 1 int value.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{Val: 0}
	cur := tmp

	for l1 != nil || l2 != nil {
		copyFromL1 := true
		switch {
		case l1 == nil:
			copyFromL1 = false
		case l2 == nil:
			// Do nothing
		case l1.Val <= l2.Val:
			// Do nothing
		default:
			copyFromL1 = false
		}

		if copyFromL1 {
			// l1から値を詰める
			cur.Next = &ListNode{Val: l1.Val}
			cur = cur.Next
			l1 = l1.Next
		} else {
			// l2から値を詰める
			cur.Next = &ListNode{Val: l2.Val}
			cur = cur.Next
			l2 = l2.Next
		}
	}

	return tmp.Next
}
