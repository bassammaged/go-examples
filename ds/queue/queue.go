package queue

type RandomQueue struct {
	items []interface{}
}

// EnQueue() appends the item at the end of the queue
func (rq *RandomQueue) EnQueue(item interface{}) {
	rq.items = append(rq.items, item)
}

// DeQueue removes the first item in the queue and return the removed item value.
func (rq *RandomQueue) DeQueue() (removedItem interface{}) {
	toRemove := rq.items[0]
	rq.items = rq.items[1:]
	return toRemove
}
