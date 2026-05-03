/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
    q := []*TreeNode{root}
	res := []int{}

	for len(q) > 0 {
		w := len(q)
		for i:= 0; i < w; i++ {
			curr := q[0]
			q = q[1:]
            if i == w-1 {
				res = append(res, curr.Val)
			}
            
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}			
		}
	}
	return res

}
