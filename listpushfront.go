package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/

func ListPushFront(l *List, data interface{}) {
	var aNodel *NodeL = &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = aNodel
		l.Tail = aNodel
	} else {
		aNodel.Next = l.Head
		l.Head = aNodel
	}
}
