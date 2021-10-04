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

func ListRemoveIf(l *List, data_ref interface{}) {
	var aPointer *NodeL = l.Head
	if l.Head != nil {
		if l.Head.Data == data_ref {
			l.Head = l.Head.Next
			aPointer.Next = l.Head
		}
		for aPointer.Next != nil {
			if aPointer.Next.Data == data_ref {
				if aPointer.Next == l.Tail {
					aPointer.Next = nil
				} else {
					aPointer.Next = aPointer.Next.Next
				}
			}
			if aPointer.Next != nil && aPointer.Next.Data != data_ref {
				aPointer = aPointer.Next
			}
		}
	}
}
