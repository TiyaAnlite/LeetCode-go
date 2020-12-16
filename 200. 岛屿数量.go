/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

type (
	Point [2]int //XY坐标点

	Queue struct {
		head   *node
		rear   *node
		length int
	}

	node struct {
		pre   *node
		next  *node
		value Point //根据实际题目的专用格式
	}
)

func main() {
	d := [][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '1'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Print(numIslands(d))
	printGrid(d)

}

func printGrid(grid [][]byte) {
	for _, d := range grid {
		fmt.Println(d)
	}
	fmt.Println("==============")
}

func numIslands(grid [][]byte) int {
	var lands int

	//统计行列信息
	gSize := [2]int{len(grid[0]), len(grid)} //X, Y

	//链式队列
	queue := NewQueue()

	//开始搜索
	for locate := scanOnce(grid); locate[0] != -1; locate = scanOnce(grid) {
		//搜索一个岛屿
		search(locate, grid, queue, gSize)
		lands++
	}
	return lands
}

func search(start Point, grid [][]byte, queue *Queue, gSize [2]int) {
	//广度优先搜索一个岛屿
	//搜索只会向后，向下进行
	//避免重复访问：每访问一个点后，删除这个点(变为水)
	//注：数组在逻辑上是按照(Y, X)格式的，极其容易混淆
	queue.Push(start)              //将第一个点推入队列
	grid[start[1]][start[0]] = '0' //先删除初始点，注意XY！！
	for status := queue.IsEmpty(); !status; status = queue.IsEmpty() {
		seek := queue.Pop() //取出一个点

		//向前搜索：X[1]
		if seek[0] > 0 && grid[seek[1]][seek[0]-1] == '1' {
			queue.Push(Point{seek[0] - 1, seek[1]})
			grid[seek[1]][seek[0]-1] = '0'
			//printGrid(grid)
		}

		//向后搜索：X[1]
		if seek[0]+1 < gSize[0] && grid[seek[1]][seek[0]+1] == '1' {
			queue.Push(Point{seek[0] + 1, seek[1]})
			grid[seek[1]][seek[0]+1] = '0'
			//printGrid(grid)
		}

		//向上搜索：Y[0]
		if seek[1] > 0 && grid[seek[1]-1][seek[0]] == '1' {
			queue.Push(Point{seek[0], seek[1] - 1})
			grid[seek[1]-1][seek[0]] = '0'
			//printGrid(grid)
		}

		//向下搜索：Y[0]
		if seek[1]+1 < gSize[1] && grid[seek[1]+1][seek[0]] == '1' {
			queue.Push(Point{seek[0], seek[1] + 1})
			grid[seek[1]+1][seek[0]] = '0'
			//printGrid(grid)
		}

	}
}

func scanOnce(grid [][]byte) Point {
	//扫描并搜索一个陆地
	//扫描只会向后，向下进行
	for y, line := range grid {
		for x, d := range line {
			if d == '1' {
				return Point{x, y}
			}
		}
	}
	//无可用陆地
	return Point{-1, -1}
}

func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

func (this *Queue) Len() int {
	return this.length
}

func (this *Queue) IsEmpty() bool {
	return this.length == 0
}

func (this *Queue) Push(o Point) {
	n := &node{nil, nil, o}
	if this.length == 0 {
		//首个节点
		this.head = n
		this.rear = this.head
	} else {
		n.pre = this.rear
		this.rear.next = n
		this.rear = n
	}
	this.length++
}

func (this *Queue) Pop() Point {
	if this.length == 0 {
		return [2]int{-1, -1}
	}
	n := this.head
	if this.head.next == nil {
		//最后一个节点
		this.head = nil
	} else {
		this.head = this.head.next
		this.head.pre.next = nil
		this.head.pre = nil
	}
	this.length--
	return n.value
}
