package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	retNode := &ListNode{Val: 0, Next: nil}
	curN1, curN2, curNode := l1, l2, retNode
	sum := 0

	for curN1 != nil || curN2 != nil {
		if curN1 != nil {
			sum += curN1.Val
			curN1 = curN1.Next
		}

		if curN2 != nil {
			sum += curN2.Val
			curN2 = curN2.Next
		}

		curNode.Next = &ListNode{Val: sum % 10}
		curNode = curNode.Next
		sum /= 10
	}

	if sum != 0 {
		curNode.Next = &ListNode{Val: sum}
	}

	return retNode.Next
}
