package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if Atoi(root.Data) > Atoi(data) {
		if root.Left != nil {
			return BTreeInsertData(root.Left, data)
		} else {
			root.Left = &TreeNode{Data: data}
		}
	} else {
		if root.Right != nil {
			return BTreeInsertData(root.Right, data)
		} else {
			root.Right = &TreeNode{Data: data}
		}
	}
	return root
}
