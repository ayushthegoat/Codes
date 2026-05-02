/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	count := 0
    return dfs(root, k, &count)
}
func dfs(root *TreeNode, k int, times *int) int {
	if root == nil {
		return -1
	}
    leftVal := dfs(root.Left, k, times)
	if leftVal != -1 {
		return leftVal
	}
    *times++
	if k == *times {
		return root.Val
	}
	return dfs(root.Right, k, times)

}
