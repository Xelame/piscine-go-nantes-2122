package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}
*/

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i < pos; i++ {
		if l.Next != nil {
			l = l.Next
		} else {
			return nil
		}
	}
	return l
}
