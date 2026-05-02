/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func maxDepth(root *TreeNode) int {
//      if root == nil {
//       return 0
//      }
//      maxL := maxDepth(root.Left)
//      maxR := maxDepth(root.Right)

//      return 1 + max(maxL, maxR)
// }
// func max(a int, b int) int {
//    if a >= b {
//       return a 
//    }
//    return b
// }

// func maxDepth(root *TreeNode) int {
//     if root ==  nil {
//       return 0
//     }
//     res:= 0
//     q := []*TreeNode{root}

//     for len(q) > 0 {
//       res++
//       width := len(q)
//       for i:=0; i < width; i++ {
//          node:= q[0]
//          q = q[1:]

//          if node.Left  != nil {
//             q = append(q, node.Left)
//          }
//          if node.Right != nil {
//             q = append(q, node.Right)
//          }
//       }
//     }
//     return res
// }

func maxDepth(root *TreeNode) int {
   if root == nil {
      return 0
   }
   lh := maxDepth(root.Left)
   rh := maxDepth(root.Right)
   return 1 + max(lh, rh)

}