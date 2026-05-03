/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
       if root.Left == nil {
		return root.Right
	   }
	   if root.Right == nil {
		return root.Left
	   }
     
		temp := findMin(root.Right)
		root.Val = temp.Val
		root.Right = deleteNode(root.Right, temp.Val)
	   

	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root 
	for cur.Left != nil {
      
		cur = cur.Left
	   
	}
	return cur
}