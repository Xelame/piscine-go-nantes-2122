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

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	var aNodel *NodeL = l.Head
	count := 0
	for aNodel != l.Tail {
		count++
		aNodel = aNodel.Next
	}
	count++
	return count
}
