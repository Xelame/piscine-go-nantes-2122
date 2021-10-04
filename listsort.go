package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}

	var currentnode *NodeI = l
	var previous *NodeI = nil
	var forward *NodeI
	for currentnode.Next != nil {
		forward = currentnode.Next
		if currentnode.Data > forward.Data {
			currentnode.Next = previous
			previous = currentnode
			currentnode = forward
		} else {
			currentnode = currentnode.Next
		}
	}
	currentnode.Next = previous
	return currentnode
}
