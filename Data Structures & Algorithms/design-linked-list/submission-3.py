class Node:

    def __init__(self, val: int, next: Optional["Node"] = None, prev: Optional["Node"] = None) -> None:
        self.val = val
        self.next = next
        self.prev = prev


class MyLinkedList:

    def __init__(self):
        self._length = 0
        self._head: Optional[Node] = None
        self._tail: Optional[Node] = None

    def get(self, index: int) -> int:
        if index < 0 or index > self._length - 1:
            return -1
        
        if index == 0:
            return self._head.val
        elif index == self._length - 1:
            return self._tail.val

        cursor = self._head
        for _ in range(index):
            cursor = cursor.next
        return cursor.val

    def addAtHead(self, val: int) -> None:
        newNode = Node(val, next=self._head)
        if self._head is None and self._tail is None:
            self._tail = newNode
        else:
            self._head.prev = newNode
        self._head = newNode
        self._length += 1

    def addAtTail(self, val: int) -> None:
        newNode = Node(val, prev=self._tail)
        if self._tail is None and self._head is None:
            self._head = newNode
        else:
            self._tail.next = newNode
        self._tail = newNode
        self._length += 1

    def addAtIndex(self, index: int, val: int) -> None:
        if index < 0 or index > self._length:
            return

        if index == 0:
            self.addAtHead(val)
        elif index == self._length:
            self.addAtTail(val)
        else:
            cursor = self._head
            for _ in range(index):
                cursor = cursor.next
            newNode = Node(val, next=cursor, prev=cursor.prev)
            cursor.prev.next = newNode
            cursor.prev = newNode
            self._length += 1

    def deleteAtIndex(self, index: int) -> None:
        if index < 0 or index > self._length - 1:
            return
        
        if self._length == 1:
            self._head = None
            self._tail = None
            self._length = 0
            return
        
        if index == 0:
            self._head.next.prev = None
            tmp = self._head
            self._head = self._head.next
            tmp.next = None
            self._length -= 1
            return
        elif index == self._length - 1:
            tmp = self._tail
            self._tail.prev.next = None
            self._tail = self._tail.prev
            tmp.prev = None
            self._length -= 1
            return
        else:
            cursor = self._head
            for _ in range(index):
                cursor = cursor.next
            cursor.prev.next = cursor.next
            cursor.next.prev = cursor.prev
            cursor.next = None
            cursor.prev = None
            self._length -= 1


# Your MyLinkedList object will be instantiated and called as such:
# obj = MyLinkedList()
# param_1 = obj.get(index)
# obj.addAtHead(val)
# obj.addAtTail(val)
# obj.addAtIndex(index,val)
# obj.deleteAtIndex(index)