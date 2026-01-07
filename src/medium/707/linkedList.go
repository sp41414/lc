package linkedList

type MyLinkedList struct {
	head *Node
}

type Node struct {
	next  *Node
	value int
}

func Constructor() MyLinkedList {
	// returns a new empty linked list struct
	return MyLinkedList{
		head: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	// case: invalid index
	if index < 0 {
		return -1
	}

	// first goes to the head of the list
	current := this.head

	// then goes through the linked list using the .next
	// while the current is not nil (out of bounds)
	// and haven't reached the index yet
	for i := 0; i < index && current != nil; i++ {
		current = current.next
	}

	// out of bounds case
	if current == nil {
		return -1
	}

	// returns the found value from the traversal above
	return current.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	// sets the head to a new node with the value
	// as the input, and the pointer to the previous head
	this.head = &Node{value: val, next: this.head}
}

func (this *MyLinkedList) AddAtTail(val int) {
	// if the linked list is empty the new tail node becomes the head
	if this.head == nil {
		this.head = &Node{value: val, next: nil}
		return
	}

	// sets the pointer to the head
	current := this.head

	// and traverses through the entire linked list
	// stopping right before we reach the actual nil tail
	// so we can set the current node's next pointer to the
	// new tail, note: could set a tail pointer but this is fine for now
	for current.next != nil {
		current = current.next
	}

	// sets the current node's next pointer to the new tail
	current.next = &Node{next: nil, value: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// case: invalid index
	if index < 0 {
		return
	}

	// case: index is 0
	if index == 0 {
		this.head = &Node{next: this.head, value: val}
		return
	}

	// traverses through the linked lists and stops right before
	// so i can't create a floating node and assign
	// the previous node's next to the new node
	current := this.head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}

	// case: the current node doesn't exist
	if current == nil {
		return
	}

	newNode := &Node{value: val, next: current.next}
	current.next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// case: invalid index
	if index < 0 {
		return
	}

	// case: the linked list is already empty
	if this.head == nil {
		return
	}

	// case: removing the head
	if index == 0 {
		this.head = this.head.next
		return
	}

	// then goes through the linked list using the .next
	// while the current is not nil (out of bounds)
	// and haven't reached the index yet (off by one, so we can set
	// the previous node's pointer to nil)
	current := this.head
	for i := 0; i < index-1 && current != nil; i++ {
		// case: index out of bounds
		if current.next == nil {
			return
		}
		current = current.next
	}

	// case: node isn't actually there to delete
	if current == nil || current.next == nil {
		return
	}

	// remove the next node by setting it to the next node's next
	current.next = current.next.next
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
