package piscine

// Returns the maximum node in the subtree started by root
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}

	return BTreeMax(root.Right)
}
