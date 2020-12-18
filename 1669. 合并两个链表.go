/*
给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。

请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。

下图中蓝色边和节点展示了操作后的结果：


请你返回结果链表的头指针。



示例 1：



输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]
解释：我们删除 list1 中第三和第四个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。
示例 2：


输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
解释：上图中蓝色的边和节点为答案链表。


提示：

3 <= list1.length <= 104
1 <= a <= b < list1.length - 1
1 <= list2.length <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-in-between-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	//注：这里注入的测试样例输出的结果似乎有问题
	list1 := []int{0, 1, 2, 3, 4, 5}
	a := 3
	b := 4
	list2 := []int{1000000, 1000001, 1000002}

	list1Node := &ListNode{0, nil}
	l1 := list1Node
	for _, v := range list1 {
		l1.Val = v
		l1.Next = &ListNode{0, nil}
		l1 = l1.Next
	}
	l1.Next = nil

	list2Node := &ListNode{0, nil}
	l2 := list2Node
	for _, v := range list2 {
		l2.Val = v
		l2.Next = &ListNode{0, nil}
		l2 = l2.Next
	}
	l2.Next = nil

	result := mergeInBetween(list1Node, a, b, list2Node)
	for ; result.Next != nil; result = result.Next {
		fmt.Println(result.Val)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	//Go当中的内存是自动管理的，故不需要手动释放节点内存
	perv, next := findNode(list1, a, b)
	perv.Next = list2
	l := list2
	for ; l.Next != nil; l = l.Next {
	} //递归到最后一个节点
	l.Next = next
	return list1
}

func findNode(list *ListNode, start int, end int) (*ListNode, *ListNode) {
	var perv *ListNode
	var next *ListNode
	for i := 0; list.Next != nil; i++ {
		if i+1 == start {
			perv = list
		}
		if i == end {
			next = list.Next
		}
		list = list.Next
	}
	return perv, next
}
