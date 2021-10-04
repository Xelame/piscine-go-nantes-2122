package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l != nil {
		var currentnode *NodeI = l
		var previous *NodeI = nil
		var forward *NodeI
		for currentnode.Next != nil {
			forward = currentnode.Next
			if currentnode.Data > forward.Data {
				currentnode.Next = previous
				previous = currentnode
				currentnode = forward
			}
		}
		currentnode.Next = previous
		return currentnode
	} else {
		return nil
	}
}
