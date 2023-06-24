package binsearch

import "fmt"

type BSearchTree struct {
	root   *Node
	length int
}

// Node
type Node struct {
	key   int
	value interface{}
	left  *Node
	right *Node
}

func (bst BSearchTree) GetLength() {
	fmt.Println(bst.length)
}

func NewBSearchTree() *BSearchTree {
	return &BSearchTree{}
}

func (bst *BSearchTree) Insert(key int, value interface{}) {
	node := &Node{key: key, value: value}

	// Add value into the root
	if bst.root == nil {
		bst.root = node
		bst.length++
	} else if key < bst.root.key { // Check the key then store based on the comparison result
		if bst.root.left == nil {
			bst.root.left = node
			bst.length++
		} else {
			fmt.Println("Recursive func() Left")
			bst.recursiveInsertNode(bst.root.left, node)
		}
	} else {
		if bst.root.right == nil {
			bst.root.right = node
			bst.length++
		} else {
			fmt.Println("Recursive func() Right")
			bst.recursiveInsertNode(bst.root.right, node)
		}
	}

}

func (bst *BSearchTree) recursiveInsertNode(currentPosition *Node, newNode *Node) {
	if currentPosition.key > newNode.key {
		if currentPosition.left == nil {
			currentPosition.left = newNode
			bst.length++
		} else {
			bst.recursiveInsertNode(currentPosition.left, newNode)
		}
	} else {
		if currentPosition.right == nil {
			currentPosition.right = newNode
			bst.length++
		} else {
			bst.recursiveInsertNode(currentPosition.right, newNode)
		}
	}
}
