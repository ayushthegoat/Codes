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
// func inorderTraversal(root *TreeNode) []int {
//     // Create an empty slice to store our result
//     // We'll collect all node values in inorder sequence here
//     res := []int{}
    
//     // Call the recursive helper function, passing:
//     // 1. The current node (starting from root)
//     // 2. A pointer to our result slice (so it can be modified directly)
//     inorder(root, &res)
    
//     // Return the filled result slice
//     return res
// }

// // inorder is a recursive helper function that traverses the tree
// // Parameters:
// //   - node: the current node being processed
// //   - res: pointer to the result slice (allows modification across recursive calls)
// // Note: We use a pointer because slices are reference types but we need to modify
// // the actual slice header (append can change length/capacity)
// func inorder(node *TreeNode, res *[]int) {
//     // BASE CASE: If the current node is nil (empty/null node)
//     // Just return without doing anything - this ends the recursion path
//     if node == nil {
//         return
//     }
    
//     // RECURSIVE STEP 1: Traverse the LEFT subtree
//     // According to inorder traversal: process left child first
//     // This will recursively visit all nodes in the left subtree
//     // and add them to res before we process the current node
//     inorder(node.Left, res)
    
//     // PROCESS CURRENT NODE: Add the current node's value to result
//     // The asterisk (*res) dereferences the pointer to access the actual slice
//     // We append the current node's value to the result slice
//     *res = append(*res, node.Val)
    
//     // RECURSIVE STEP 2: Traverse the RIGHT subtree
//     // After processing left subtree and current node, process right subtree
//     // This will recursively visit all nodes in the right subtree
//     // and add them to res after the current node
//     inorder(node.Right, res)
// }

// Tree structure:
//     2
//    / \
//   1   3

// Execution flow:
// 1. inorderTraversal(root=2) called
//    - res = [] (empty slice)
   
// 2. inorder(2, &res) called
//    - node=2 is not nil
//    - Call inorder(2.Left=1, &res)
   
// 3. inorder(1, &res) called
//    - node=1 is not nil
//    - Call inorder(1.Left=nil, &res)
   
// 4. inorder(nil, &res) called
//    - node=nil, so return (base case)
   
// 5. Back to inorder(1, &res)
//    - Append node 1's value: *res = append(*res, 1) -> res = [1]
//    - Call inorder(1.Right=nil, &res)
   
// 6. inorder(nil, &res) called - returns immediately
   
// 7. Back to inorder(2, &res)
//    - Append node 2's value: *res = append(*res, 2) -> res = [1, 2]
//    - Call inorder(2.Right=3, &res)
   
// 8. inorder(3, &res) called
//    - Append node 3's value: *res = append(*res, 3) -> res = [1, 2, 3]
   
// 9. Back to inorderTraversal - returns [1, 2, 3]




func inorderTraversal(root *TreeNode) []int {
    res := []int{}        // Stores our final result
    stack := []*TreeNode{} // Explicit stack to track nodes (replaces recursion call stack)
    cur := root           // Current node we're processing

    // Continue while we have nodes to process (either current node exists OR stack has nodes)
    for cur != nil || len(stack) > 0 {
        
        // STEP 1: Go as far LEFT as possible
        // Push all left children onto the stack
        for cur != nil {
            stack = append(stack, cur)  // Push current node to stack
            cur = cur.Left              // Move to left child
        }
        // After this loop, cur == nil (reached the bottom-left)
        
        // STEP 2: Process a node (when we can't go left anymore)
        cur = stack[len(stack)-1]  // Pop from stack (get top element)
        stack = stack[:len(stack)-1] // Remove it from stack
        
        // STEP 3: Visit/process the current node
        res = append(res, cur.Val)  // Add value to result
        
        // STEP 4: Go RIGHT (and then the outer loop will push all left children of this right node)
        cur = cur.Right
    }
    
    return res
}