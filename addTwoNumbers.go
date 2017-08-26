
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret ListNode

	cur := &ret
	var lastOver int
	for {
		cur.Val, lastOver = lastOver, 0
		if l1 != nil {
			cur.Val += l1.Val
		}
		if l2 != nil {
			cur.Val += l2.Val
		}
		if cur.Val > 9 {
			cur.Val %= 10
			lastOver = 1
		}

		if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
			if lastOver > 0 {
				cur.Next = new(ListNode)
				cur.Next.Val = lastOver
			}
			break
		} else {
			if l1 != nil {
				l1 = l1.Next
			}
			if l2 != nil {
				l2 = l2.Next
			}
			if l1 == nil && l2 == nil {
				break
			}
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	return &ret
}
