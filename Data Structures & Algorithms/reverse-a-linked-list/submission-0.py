# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        reversedList = ListNode(val=head.val)
        head = head.next
        while head is not None:
            temp = head
            head = head.next
            temp.next = reversedList
            reversedList = temp
        return reversedList