/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if abs(lh - rh) > 1 {
		return false
	}
	lb := isBalanced(root.Left)
	rb := isBalanced(root.Right)
	return lb && rb
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)
	return 1 + max(lh, rh)
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}