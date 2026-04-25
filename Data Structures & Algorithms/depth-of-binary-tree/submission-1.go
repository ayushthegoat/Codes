/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
     if root == nil {
      return 0
     }
     maxL := maxDepth(root.Left)
     maxR := maxDepth(root.Right)

     return 1 + max(maxL, maxR)
}
func max(a int, b int) int {
   if a >= b {
      return a 
   }
   return b
}