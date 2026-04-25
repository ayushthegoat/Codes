/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int        // Value stored in the node
 *     Left *TreeNode // Pointer to left child node
 *     Right *TreeNode // Pointer to right child node
 * }
 */

// inorderTraversal is the main function that takes the root of a binary tree
// and returns a slice containing the values in inorder traversal order
// (Left -> Node -> Right)
func inorderTraversal(root *TreeNode) []int {
    // Create an empty slice to store our result
    // We'll collect all node values in inorder sequence here
    res := []int{}
    
    // Call the recursive helper function, passing:
    // 1. The current node (starting from root)
    // 2. A pointer to our result slice (so it can be modified directly)
    inorder(root, &res)
    
    // Return the filled result slice
    return res
}

// inorder is a recursive helper function that traverses the tree
// Parameters:
//   - node: the current node being processed
//   - res: pointer to the result slice (allows modification across recursive calls)
// Note: We use a pointer because slices are reference types but we need to modify
// the actual slice header (append can change length/capacity)
func inorder(node *TreeNode, res *[]int) {
    // BASE CASE: If the current node is nil (empty/null node)
    // Just return without doing anything - this ends the recursion path
    if node == nil {
        return
    }
    
    // RECURSIVE STEP 1: Traverse the LEFT subtree
    // According to inorder traversal: process left child first
    // This will recursively visit all nodes in the left subtree
    // and add them to res before we process the current node
    inorder(node.Left, res)
    
    // PROCESS CURRENT NODE: Add the current node's value to result
    // The asterisk (*res) dereferences the pointer to access the actual slice
    // We append the current node's value to the result slice
    *res = append(*res, node.Val)
    
    // RECURSIVE STEP 2: Traverse the RIGHT subtree
    // After processing left subtree and current node, process right subtree
    // This will recursively visit all nodes in the right subtree
    // and add them to res after the current node
    inorder(node.Right, res)
}