package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	// Create a nodel
	var aNodel *NodeL = &NodeL{Data: data, Next: nil}
	// If list is empty
	if l.Head == nil {
		l.Head = aNodel
		l.Tail = aNodel
	} else {
		// Else add the nodel at the end
		l.Head.Next = aNodel
		// Update the end list pointer
		l.Head = aNodel
	}
}
