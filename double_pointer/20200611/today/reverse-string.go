/*
344. 反转字符串

编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
	输入：["h","e","l","l","o"]
	输出：["o","l","l","e","h"]

示例 2：
	输入：["H","a","n","n","a","h"]
	输出：["h","a","n","n","a","H"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

/*
找到中间节点的位置
循环，将第一个和最有一个交换
*/
func reverseString(s []byte) {
	length := len(s)
	if length <= 1 {
		return
	}

	middle := length / 2

	for i := 0; i < middle; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// 另一种双指针方法
// 定义left:0, right:length-1指针
// 在 left < right的时候，一次交换
// 交换后右移left， 左移right
func reverseString2(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}