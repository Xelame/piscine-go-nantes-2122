package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var currentnode *NodeI = l
	var previous *NodeI = nil
	var forward *NodeI
	for currentnode.Next != nil && currentnode != nil {
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
