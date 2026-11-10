/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	
	// Initialize
	var cursor *ListNode
	if list1.Val < list2.Val {
		cursor = list1
		list1 = list1.Next
	} else {
		cursor = list2
		list2 = list2.Next
	}

	mergedList := cursor

	for list1 != nil || list2 != nil {
		if list1 == nil {
			cursor.Next = list2
			return mergedList
		} else if list2 == nil {
			cursor.Next = list1
			return mergedList
		}

		if list1.Val < list2.Val {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}
		cursor = cursor.Next
	}
	return mergedList
}
