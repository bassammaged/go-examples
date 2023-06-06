package queue

type RandomQueue struct {
	items []interface{}
}

func (rq *RandomQueue) EnQueue(item interface{}) {
	rq.items = append(rq.items, item)
}

func (rq *RandomQueue) DeQueue() (removedItem interface{}) {
	toRemove := rq.items[0]
	rq.items = rq.items[1:]
	return toRemove

}
