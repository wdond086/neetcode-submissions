type MyLinkedList struct {
    length int
	head *Node
	tail *Node
}

type Node struct {
	val int
	next *Node
	prev *Node
}


func Constructor() MyLinkedList {
    return MyLinkedList{length: 0}
}


func (this *MyLinkedList) Get(index int) int {
    if index < -1 || index > this.length - 1 {
		return -1
	}

	if index == 0 {
		return this.head.val
	} else if index == this.length - 1 || index == -1 {
		return this.tail.val
	}

	cursor := this.head
	for range index {
		cursor = cursor.next
	}
	return cursor.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newNode := &Node{val: val, next: this.head}
	if this.head == nil && this.tail == nil {
		this.tail = newNode
	} else {
		this.head.prev = newNode
	}
	this.head = newNode
	this.length += 1
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &Node{val: val, prev: this.tail}
	if this.tail == nil && this.head == nil {
		this.head = newNode
	} else { // List was empty
		this.tail.next = newNode
	}
	this.tail = newNode
	this.length += 1
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else {
		cursor := this.head
		for range index {
			cursor = cursor.next
		}
		newNode := &Node{val: val, next: cursor, prev: cursor.prev}
		cursor.prev.next = newNode
		cursor.prev = newNode
		this.length += 1
	}
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index > this.length - 1 {
		return
	}

	if this.length == 1 {
		this.head = nil
		this.tail = nil
		this.length = 0
		return
	}

	if index == 0 {
		this.head.next.prev = nil
		tmp := this.head
		this.head = this.head.next
		tmp.next = nil
		this.length -= 1
		return
	} else if index == this.length - 1 {
		tmp := this.tail
		this.tail.prev.next = nil
		this.tail = this.tail.prev
		tmp.prev = nil
		this.length -= 1
		return
	} else {
		cursor := this.head
		for range index {
			cursor = cursor.next
		}
		cursor.prev.next = cursor.next
		cursor.next.prev = cursor.prev
		cursor.next = nil
		cursor.prev = nil
		this.length -= 1
	}
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */