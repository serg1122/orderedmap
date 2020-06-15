package orderedmap

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func (n Node) HasPrev() bool {

	return n.prev != nil
}

func (n *Node) SetPrev(prev *Node) {

	n.prev = prev
}

func (n Node) Prev() *Node {

	return n.prev
}

func (n Node) HasNext() bool {

	return n.next != nil
}

func (n *Node) SetNext(next *Node) {

	n.next = next
}

func (n *Node) Next() *Node {

	return n.next
}

func (n Node) Value() interface{} {

	return n.value
}

func (n *Node) SetValue(value interface{}) {

	n.value = value
}
