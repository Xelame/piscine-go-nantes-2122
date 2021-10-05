package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)
		f(root.Data)
	}
}