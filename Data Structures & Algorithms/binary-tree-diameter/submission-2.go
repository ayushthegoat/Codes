/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func diameterOfBinaryTree(root *TreeNode) int {
//     if root == nil {
//         return 0  
//     }
    
//     // Path through root
//     throughRoot := maxDepth(root.Left) + maxDepth(root.Right)
    
//     // Diameter in left subtree (might be larger)
//     leftDiameter := diameterOfBinaryTree(root.Left)
    
//     // Diameter in right subtree (might be larger)
//     rightDiameter := diameterOfBinaryTree(root.Right)
    
//     // Return the maximum of all three
//     return max(throughRoot, max(leftDiameter, rightDiameter))
// }
// func maxDepth(root *TreeNode) int {
//    if root == nil {
//       return 0
//    }
//    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
// }
// func diameterOfBinaryTree(root *TreeNode) int {
//    if root == nil {
//     return 0
//    }
//    d1 := depth(root.Left) + depth(root.Right)
//    d2 := diameterOfBinaryTree(root.Left)
//    d3 := diameterOfBinaryTree(root.Right)
//    return max(d1, max(d2, d3))



// }
// func depth (root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     lh := depth(root.Left)
//     rh := depth(root.Right)
//     return 1 + max(lh, rh)
// }
var d int 
func diameterOfBinaryTree(root *TreeNode) int {
    d = 0
    if root == nil {
        return 0
    }
    height(root)
    return d
}

func height(root *TreeNode) int {
   if root == nil {
    return 0
   }
   lh := height(root.Left)
   rh := height(root.Right)
   curH := lh + rh
   if curH > d {
    d = curH
   }
   return 1 + max(lh, rh)
}