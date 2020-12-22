/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	/*
	    3A
	   / \
	  9B  20C
	    /  \
	   15D   7E
	*/
	tree_c := &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	tree_b := &TreeNode{Val: 9}
	root_a := &TreeNode{Val: 3, Left: tree_b, Right: tree_c}
	fmt.Println(zigzagLevelOrder(root_a))
	fmt.Println(zigzagLevelOrder(nil))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var direction int
	pDirection := &direction  //方向指示，对2取模
	output := [][]int{}       //输出数组
	heap := []*TreeNode{root} //栈遍历数组
	*pDirection += 1
	output = search103(output, heap, pDirection)
	return output
}

func search103(output [][]int, heap []*TreeNode, pDirection *int) [][]int {
	/*
		按层读取节点并输出，即广度优先
		思路：按照栈的顺序读取值是反向的，做到每行都可以翻转方向输出
		但依旧会根据状态变量确定方向，确定子节点的访问顺序(从左或丛右)
		所有当前层节点在本函数体内只需要被访问一次，效率高，无需中途调转方向
		预先创建的栈数组提高索引效率，并减少新建或动态扩增数组时的开销
		该方法虽然效率高，但是每一层都需要预先创建当前深度最大容量数组便于快速索引，树的深度大时占用内存高
	*/
	layer := make([]int, len(heap))             //当前层节点
	next_heap := make([]*TreeNode, len(heap)*2) //复数层栈容量
	var len_next_heap int
	for i := len(heap) - 1; i >= 0; i-- {
		//弹栈，读取值
		layer[i] = heap[i].Val
		//确定遍历子节点方向
		if *pDirection%2 == 0 {
			//从左往右
			len_next_heap += readTreeNode(heap[i].Left, next_heap, len_next_heap)
			len_next_heap += readTreeNode(heap[i].Right, next_heap, len_next_heap)
		} else {
			//从右往左
			len_next_heap += readTreeNode(heap[i].Right, next_heap, len_next_heap)
			len_next_heap += readTreeNode(heap[i].Left, next_heap, len_next_heap)
		}
	}
	next_heap = next_heap[:len_next_heap] //裁剪栈
	output = append(output, layer)
	if len(next_heap) > 0 {
		//下一层递归
		*pDirection += 1
		output = search103(output, next_heap, pDirection)
	}
	return output
}

func readTreeNode(n *TreeNode, heap []*TreeNode, heapIndex int) int {
	//读取子节点，如果有子节点，入栈并返回1，否则返回0
	if n == nil {
		return 0
	} else {
		heap[heapIndex] = n
		return 1
	}
}
