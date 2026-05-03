/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := [][]int{}
    q := []*TreeNode{root}
    levelT := 1
    for len(q) > 0 {
        
        w := len(q)
        level := []int{}
        for i := 0; i < w; i++ {
            cur := q[0]
            q = q[1:]
        
            if levelT % 2 != 0 {
                level = append(level, cur.Val)
            } else {
                level = append([]int{cur.Val}, level...)
            }

            if cur.Left != nil { 
                q = append(q, cur.Left) 
            }
            if cur.Right != nil { 
                q = append(q, cur.Right) 
            }
        }
        levelT++
        res = append(res, level)
    }
    return res
}
