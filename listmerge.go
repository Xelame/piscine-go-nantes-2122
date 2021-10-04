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
	if l1.Head != nil && l2.Head != nil {
		l1.Tail.Next = l2.Head
	}
}
