/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
//     // AGAR ROOT HI NULL HAI → MATLAB TREE KHAALI HAI
//     // TOH LCA BHI NULL HI HOGA, WAPAS NULL BHJAA DO
//     if root == nil {
//         return root
//     }
    
//     // CURRENT NODE KI VALUE NIKAL LO (COMPARE KARNE KE LIYE)
//     valOfCur := root.Val
    
//     // CASE 1: DONO NODES (p aur q) CURRENT NODE SE BADE HAIN
//     // MATLAB DONO RIGHT SUBTREE MEIN HONGE
//     // ISLIYE RIGHT MEIN RECURSION KARO
//     if p.Val > valOfCur && q.Val > valOfCur {
//         return lowestCommonAncestor(root.Right, p, q)
//     } 
    
//     // CASE 2: DONO NODES (p aur q) CURRENT NODE SE CHOTE HAIN
//     // MATLAB DONO LEFT SUBTREE MEIN HONGE
//     // ISLIYE LEFT MEIN RECURSION KARO
//     else if p.Val < valOfCur && q.Val < valOfCur {
//         return lowestCommonAncestor(root.Left, p, q)
//     }
    
//     // CASE 3: EK NODE CHOTA HAI, EK BADA HAI
//     // YA PHIR KOI EK NODE CURRENT NODE KE EQUAL HAI
//     // ISKA MATLAB CURRENT NODE HI UNKA LCA HAI
//     // ISLIYE CURRENT NODE KO RETURN KAR DO
//     return root
// }

// Tree Structure:
//         5  ←─── root (valOfCur = 5)
//        / \
//       3   8
//      / \   \
//     1   4   9

// Scenario: p=1, q=4 (don't chote hai root se)

// Step 1: root = 5
//         p.Val=1 > 5? NO
//         p.Val=1 < 5? YES
//         q.Val=4 > 5? NO
//         q.Val=4 < 5? YES
//         → ELSE IF condition TRUE
//         → LEFT MEIN JAAO (lowestCommonAncestor(root.Left, p, q))

// Step 2: root = 3
//         p.Val=1 > 3? NO
//         p.Val=1 < 3? YES
//         q.Val=4 > 3? YES  (ab ek node chota, ek bada)
//         → ELSE IF condition FALSE
//         → ELSE (return root) → return 3 ✅

// Answer: 3 (Lowest Common Ancestor)


func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    //itrative apparoch 
    cur := root



   //we directly iterate and then keep on checking if that 
   //values is less or right if the values is actually less we move left
   //if the values are larger then we move right
   //and then when there is no wehere to move when teh code falls in the 
   //else block we have find the actually LCA 
   //we returner it
    for cur != nil {
        if p.Val > cur.Val && q.Val > cur.Val {
            cur = cur.Right
        } else if p.Val < cur.Val && q.Val < cur.Val {
            cur = cur.Left
        } else {
            return cur
        }
    }
    return root

}