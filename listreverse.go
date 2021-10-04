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

// A commenter
func ListReverse(l *List) {
	var currentnode *NodeL = l.Head
	var previous *NodeL = nil
	var forward *NodeL
	l.Head, l.Tail = l.Tail, l.Head
	if l.Head != nil {
		for currentnode != nil {
			forward = currentnode.Next
			currentnode.Next = previous
			previous = currentnode
			currentnode = forward
		}
	}
}
