/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2


提示：你可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	//只包含小写字母，码表97 - 122
	counter := make([]int, 26)
	//第一次，扫描统计
	for _, char := range s {
		counter[(char-97)%26] += 1
	}
	//第二次，找到第一个
	for i, char := range s {
		if counter[(char-97)%26] == 1 {
			return i
		}
	}
	return -1
}
