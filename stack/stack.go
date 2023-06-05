package stack

type RandomStack struct {
	items []interface{}
}

// Push() will append the item in the last index
func (rs *RandomStack) Push(item interface{}) {
	rs.items = append(rs.items, item)
}

// Pop() will remove the last item in the array and return the removed item
func (rs *RandomStack) Pop() (item interface{}) {
	length := len(rs.items) - 1
	toRemove := rs.items[length]
	rs.items = rs.items[:length]
	return toRemove
}
