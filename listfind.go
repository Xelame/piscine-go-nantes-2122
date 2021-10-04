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

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	var result *interface{}
	for l.Head != nil {
		if comp(l.Head.Data, ref) {
			result = &l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return result
}
