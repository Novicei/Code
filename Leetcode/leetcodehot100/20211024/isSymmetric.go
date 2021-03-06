package _0211024

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil || right == nil) || (left.Val != right.Val) {
		return false
	}
	return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
