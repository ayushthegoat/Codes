/**
 * Definition for binary tree node
 * Har node ke paas ek value hai, left child hai, right child hai
 */

// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
//     // case 1: dono trees khaali hain
//     // matlab main tree bhi khaali, subtree bhi khaali
//     // toh subtree match ho gaya -> true return karo
//     if root == nil && subRoot == nil {
//         return true
//     }
    
//     // case 2: main tree khaali hai, par subtree khaali nahi hai
//     // jaise main tree mein kuch hai hi nahi, toh subtree kahan milega?
//     // isliye false return karo
//     if root == nil && subRoot != nil {
//         return false
//     }
    
//     // case 3: main tree mein nodes hain, par subtree khaali hai
//     // generally khaali subtree har tree ka subtree hota hai
//     // but is problem mein subRoot kabhi nil nahi aata
//     // phir bhi safe rahne ke liye false return karo
//     if root != nil && subRoot == nil {
//         return false
//     }
    
//     // ab current node check karo: kya ye subtree ke root ke barabar hai?
//     // agar haan, toh true return karo, khela khatam
//     if isSameTree(root, subRoot) {
//         return true
//     }
    
//     // agar current node match nahi karta, toh do jagah dhundho:
//     // 1. left subtree ke andar
//     // 2. right subtree ke andar
//     // kahin bhi mil gaya toh true, nahi toh false
//     return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
// }

// func isSameTree(p *TreeNode, q *TreeNode) bool {
//     // case 1: dono nodes nil hain
//     // matlab yeh branch khatam ho gayi, aur dono taraf same hai
//     // isliye true return karo
//     if p == nil && q == nil {
//         return true
//     }
    
//     // case 2: pehla node nil hai, doosra nil nahi hai
//     // ek taraf branch khatam ho gayi, doosri taraf abhi bhi nodes hain
//     // structure match nahi karta, isliye false
//     if p == nil && q != nil {
//         return false
//     }
    
//     // case 3: pehla node nil nahi hai, doosra nil hai
//     // same as above, structure mismatch
//     if p != nil && q == nil {
//         return false
//     }
    
//     // case 4: dono nodes nil nahi hain
//     // ab values compare karo
//     if p != nil && q != nil {
//         // agar values alag hain, toh trees equal nahi ho sakte
//         if p.Val != q.Val {
//             return false
//         }
//     }
    
//     // agar values match kar gaye, toh ab left subtree check karo
//     // aur right subtree check karo
//     // dono sides match karni chahiye, isliye AND (&&) laga hai
//     return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }

// maan lo yeh tree hai:

// root (main tree):
//         5
//        / \
//       3   7
//      / \   \
//     1   4   9

// subRoot (jo dhundo) :
//         3
//        / \
//       1   4

// ab dekhte hai kya hota hai:

// pehla call: isSubtree(root=5, subRoot=3)
//   - kya 5 ke andar 3 subtree hai?
//   - pehle check karo: kya 5 aur 3 same tree hai? nahi
//   - toh left mein dhundo: isSubtree(3, 3)
//   - right mein dhundo: isSubtree(7, 3)

// ab left call: isSubtree(3, 3)
//   - kya 3 ke andar 3 subtree hai?
//   - pehle check karo: isSameTree(3, 3)
//     * 3 == 3 (value match)
//     * left mein 1==1 (match)
//     * right mein 4==4 (match)
//     * isliye true return karega
//   - toh isSubtree(3,3) true return kar diya

// ab wapas aao: isSubtree(5,3) ko true mil gaya
// isliye final answer = true



//optimizattion serialize and check contains 
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return subRoot == nil
    }
    rootSer := serialize(root)
    subRootSer := serialize(subRoot)
    return strings.Contains(rootSer, subRootSer)
}
func serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    return fmt.Sprintf("%d%s%s", root.Val, serialize(root.Left), serialize(root.Right))
}