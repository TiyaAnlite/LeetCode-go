/*
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

例如，

给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2
你应该返回如下子树:

      2
     / \
    1   3
在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	/*
	    4A
	   / \
	  2B  7C
	    /  \
	   5D   8E
	*/
	tree_c := &TreeNode{Val: 7, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 8}}
	tree_b := &TreeNode{Val: 2}
	root_a := &TreeNode{Val: 4, Left: tree_b, Right: tree_c}
	fmt.Println(searchBST(root_a, 7))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

func searchBST(root *TreeNode, val int) *TreeNode {
	resault := root
	for resault != nil {
		if val == resault.Val {
			return resault
		} else if val < resault.Val {
			resault = resault.Left
		} else {
			resault = resault.Right
		}
	}
	return resault
}
