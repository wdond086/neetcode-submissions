

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	reversedList := &ListNode{Val: head.Val, Next: nil}
	head = head.Next
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = reversedList
		reversedList = temp
	}
	return reversedList
}
