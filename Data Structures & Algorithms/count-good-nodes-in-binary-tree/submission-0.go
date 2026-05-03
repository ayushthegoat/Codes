/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return solve(root, math.MinInt64)
}
func solve(root *TreeNode, max int) int {
    if root == nil {
		return 0
	}
	count := 0
    if root.Val >= max {
		max = root.Val
		count++
	} 
    ls := solve(root.Left, max)
	rs := solve(root.Right, max)
	return count + ls + rs  
}
