package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var currentnode *NodeI = l
	var previous *NodeI = nil
	var forward *NodeI
	if l != nil {
		for currentnode.Next != nil {
			forward = currentnode.Next
			if currentnode.Data > forward.Data {
				currentnode.Next = previous
				previous = currentnode
				currentnode = forward
			}
		}
	}
	currentnode.Next = previous
	return currentnode
}
