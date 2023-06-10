package linkedlist

import (
	"errors"
	"fmt"
)

// Create a single linkedlist struct
type SingleLinkedlist struct {
	head   *Node
	length int
}

// Create a Node struct
type Node struct {
	data interface{}
	next *Node
}

// Empty Single linkedlist struct Provision
func NewSingleLinkedlist() *SingleLinkedlist {
	return &SingleLinkedlist{}
}

// InsertAtBegin function inserts the value at the start and push the current head to the next node
func (sl *SingleLinkedlist) InsertAtBegin(item interface{}) {

	node := &Node{
		data: item,
	}

	// Insert the value at the linkedlist head
	if sl.insertAtHead(node) != nil {
		// Insert the value at the start and push the head on node forward
		currentHead := sl.head
		sl.head = node
		sl.head.next = currentHead
	}

	sl.length++
}

// insertAtHead function adds the value to the single linkedlist's head
// It returns an error if the head is already occupied.
func (sl *SingleLinkedlist) insertAtHead(node *Node) error {
	if sl.head == nil {
		sl.head = node
		return nil
	}
	return errors.New("single linkedlist head is not empty")
}

func (sl *SingleLinkedlist) InsertAtEnd(item interface{}) {
	node := &Node{
		data: item,
	}

	// Insert the value at the linkedlist head
	if sl.insertAtHead(node) != nil {
		// iterate over the linkedlist and get the last item
		currentNode := sl.head // Tracking the current node
		for i := 1; i <= sl.length; i++ {
			if currentNode.next != nil {
				currentNode = currentNode.next
			} else {
				currentNode.next = node
			}
		}
	}

	sl.length++
}

// Traversal() iterates over the single linkedlist and print out the value of each node.
// It returns error if the linkedlist has no value at all.
func (sl SingleLinkedlist) Traversal() error {
	// Check if the head is empty
	if sl.head == nil || sl.length < 1 {
		return errors.New("the single linkedlist is empty")
	}

	node := sl.head
	for i := 1; i <= sl.length; i++ {
		fmt.Println(node.data)
		node = node.next
	}

	return nil
}

// Update function updates the Node.data value with a new value
// It returns error if the value doesn't exist
func (sl *SingleLinkedlist) Update(item interface{}, replaceWith interface{}) error {
	currentNode := sl.head
	for i := 1; i <= sl.length; i++ {
		if currentNode.data == item {
			currentNode.data = replaceWith
			return nil
		}
		currentNode = currentNode.next
	}
	return errors.New("the value doesn't exist")
}

func (sl *SingleLinkedlist) Delete(item interface{}) error {
	// Incase the linkedlist is empty
	if sl.length == 0 {
		return errors.New("linkedlist has no values")
	}

	currentNode := sl.head
	// Check the head value and change it if the value equal to the item
	if currentNode.data == item {
		sl.head = currentNode.next
		sl.length--
		return nil
	}

	previousNode := currentNode
	for i := 1; i < sl.length; i++ {
		if previousNode.next.data == item {
			previousNode.next = previousNode.next.next
			sl.length--
			return nil
		} else {
			previousNode = previousNode.next
		}
	}

	return errors.New("the value doesn't exist")
}
