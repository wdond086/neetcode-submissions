# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if list1 is None:
            return list2
        elif list2 is None:
            return list1
        
        # Initialize the mergedList
        cursor: ListNode
        if list1.val < list2.val:
            cursor = list1
            list1 = list1.next
        else:
            cursor = list2
            list2 = list2.next

        mergedList = cursor

        while list1 is not None or list2 is not None:
            if list1 is None:
                cursor.next = list2
                return mergedList
            elif list2 is None:
                cursor.next = list1
                return mergedList

            if list1.val < list2.val:
                cursor.next = list1
                list1 = list1.next
            else:
                cursor.next = list2
                list2 = list2.next
            cursor = cursor.next
        return mergedList
