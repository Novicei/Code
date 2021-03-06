package _0210912

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = BuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
