/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func insertIntoBST(root *TreeNode, val int) *TreeNode {
//     // CASE 1: Agar tree khaali hai (root nil hai)
//     // Matlab abhi tak koi node exist nahi karta
//     if root == nil {
//         // Naya node banao with given value
//         // & ka matlab address (pointer) return karna
//         root = &TreeNode{  
//             Val: val,
//             Left: nil,   // Left mein abhi kuch nahi
//             Right: nil,  // Right mein abhi kuch nahi
//         }
//         return root  // Naya bana node return karo
//     }
    
//     // CASE 2: Tree khaali nahi hai, toh sahi jagah dhundo
    
//     // Agar naya value current node se BADA hai
//     // Toh yeh RIGHT subtree mein jaayega
//     if val > root.Val {
//         // Right child mein insert karo, aur jo naya subtree bane
//         // usse root.Right mein assign karo
//         root.Right = insertIntoBST(root.Right, val)
//     } else if val < root.Val {
//         // Left child mein insert karo
//         root.Left = insertIntoBST(root.Left, val)
//     }
    
//     // Agar val EQUAL hai root.Val ke? 
//     // BST mein typically duplicate values allow nahi hote
//     // Toh kuch mat karo, wapas root return karo
    
//     return root  // Root ko waise hi return karo (structure same)
// }
func insertIntoBST(root *TreeNode, val int) *TreeNode {
   if root == nil {
    return &TreeNode{
        Val: val,
        Left: nil,
        Right: nil,
    }
    return root
   }
   cur := root 
   for cur != nil {
     if val < cur.Val {
        //if the value is less that the curernt node then ofcourse
        //the left node is empty the according to the BST property 
        //we knwo that we have found the spot for the node
        //we diredctly attach and return otherwise we keep iterating
          if cur.Left == nil {
            cur.Left = &TreeNode{
                Val: val,
                Left: nil,
                Right: nil,
            }
            //as soon as we attach we return otherwise the loop wil be infinite
            return root
          }
          cur = cur.Left
     } else {
        if cur.Right == nil {
            cur.Right = &TreeNode{
                Val: val,
                Left: nil,
                Right: nil,
            }
            return root
        }
        cur = cur.Right
     }
   }
   return root

}