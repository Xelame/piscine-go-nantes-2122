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

func ListMerge(l1 *List, l2 *List) {
	// Test if each list isn't empty
	if l1.Head != nil && l2.Head != nil {
		l1.Tail.Next = l2.Head
		// If the 1st list is, copy the 2nd on the 1st
	} else if l2.Head != nil {
		l1.Head = l2.Head
		// Same principle but with the 2nd list
	} else {
		l2.Head = l1.Head
	}
}
