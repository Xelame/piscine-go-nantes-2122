package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}

	return BTreeMin(root.Left)
}
