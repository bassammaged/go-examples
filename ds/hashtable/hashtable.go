package hashtable

import (
	"errors"
)

const ArraySize int = 20

// Hash table struct
type HashTable struct {
	Array [ArraySize]*bucket
}

// Bucket (Singly linkedlist) structure
type bucket struct {
	head   *bucketNode
	length int
}

// Bucket node
type bucketNode struct {
	id   int
	name string
	next *bucketNode
}

// Create an empty hash table struct
// return hash table pointer
func NewHashTable() *HashTable {
	h := HashTable{}
	for i := range h.Array {
		h.Array[i] = &bucket{}
	}
	return &h
}

// hash() hashing the word into an index value.
// Argument: word `index`
// Return index `int`
func Hash(word string) (index int) {
	var sum int
	for _, char := range word {
		sum += int(char)
	}
	return sum % ArraySize
}

// Insert() adds the values into the hash table
func (h *HashTable) Insert(id int, name string) {
	bucket := h.getBucket(name)
	bucket.InsertAtBegin(id, name)
}

// getBucket() It's not exported function, allow to hash the word and get the corresponding bucket
func (h HashTable) getBucket(name string) *bucket {
	hashIndex := Hash(name)
	return h.Array[hashIndex]
}

// InsertAtBegin() inserts a new node at the beginning of the bucket
func (b *bucket) InsertAtBegin(id int, name string) {
	bNode := &bucketNode{id: id, name: name}

	// If the head is empty
	if b.head == nil {
		b.head = bNode
	} else {
		perviousHead := b.head
		b.head = bNode
		b.head.next = perviousHead
	}
	b.length++
}

// Search() allows search for specific value in the Hashtable
func (h HashTable) Search(word string) error {
	bucket := h.getBucket(word)

	// Check if the bucket length is zero
	if bucket.length == 0 {
		return errors.New("the bucket has no value")
	}
	// TODO: search in bucket (linkedlist) and print out the result
	// OR return an error if the value doesn't exist in the bucket
	return nil
}
