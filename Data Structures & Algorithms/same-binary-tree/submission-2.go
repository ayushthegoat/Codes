/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	//yaha pe chek kro if null
//     if p == nil && q == nil{
// 		return true
// 	}
// 	//yaha pe check krow ki bhia if one not null and one null
// 	if p != nil && q == nil {
// 		return false
// 	}
// 	// yaha pe if opther wala if not null then 
// 	if p == nil && q != nil {
// 		return false
// 	}
// 	// yaha pe bhai yaar agar null nahi to wow we then compare the value and then move forward 
// 	if p != nil && q != nil {
// 		if p.Val != q.Val {
// 			return false
// 		} 
// 	}
// 	//ye return true tab rehta jab ye tree mein gauranteed only root nodes hote but huamre pass to like multiple nodes hai right
// 	// so kyu na yehi same fucntio mein recursion mein for the left and right subtree call kardalu right ??????
// 	// that is what we do in the below isSameTree recusrion call right for both the nodes using he left and right nodes in each call
// 	//return true
// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }
func isSameTree(p *TreeNode, q *TreeNode) bool {
	q1 := []*TreeNode{p}
	q2 := []*TreeNode{q}

	for len(q1) > 0 && len(q2) > 0 {
		w := len(q1)
		for i:= 0; i < w; i++ {
			np := q1[0]
			nq := q2[0]
			q1 = q1[1:]
			q2 = q2[1:]

			if np == nil && nq == nil {
				continue
			}
			if np == nil || nq == nil || np.Val != nq.Val {
				return false
			}
			q1 = append(q1, np.Left, np.Right)
			q2 = append(q2, nq.Left, nq.Right)
		}
	}
	return len(q1) == 0 && len(q2) == 0
}


