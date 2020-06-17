/*
面试题 17.18. 最短超串
假设你有两个数组，一个长一个短，短的元素均不相同。找到长数组中包含短数组所有的元素的最短子数组，其出现顺序无关紧要。
返回最短子数组的左端点和右端点，如有多个满足条件的子数组，返回左端点最小的一个。若不存在，返回空数组。

示例 1:
	输入:
		big = [7,5,9,0,2,1,3,5,7,9,1,1,5,8,8,9,7]
		small = [1,5,9]
	输出: [7,10]
示例 2:
	输入:
		big = [1,2,3]
		small = [4]
	输出: []
提示：
	big.length <= 100000
	1 <= small.length <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-supersequence-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


【思路】

1, 将small数组中的数存在map中，其value初始化为-1
2, 遍历big数组，map存储small数组中每一个值在big数组中的位置，并更新
3, 当在big数组中找齐了所有在small数组中的数字后，就用当前下标i减去map中value的最小值（即位置的最小值）
	得到的差即为“包含短数组所有的元素的子数组长度”

时间复杂度为O(nm)

作者：yuanninesuns
链接：https://leetcode-cn.com/problems/shortest-supersequence-lcci/solution/chao-xiang-xi-jie-fa-yi-ding-neng-kan-dong-by-yuan/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
package main

import "math"

func shortestSeq(big []int, small []int) []int {
	var result []int
	bigLen, smallLen := len(big), len(small)
	if bigLen == 0 || smallLen == 0 || bigLen < smallLen {
		return result
	}

	// 初始化，假设最大值
	result = append(result, 0, bigLen)

	// 将small数据放到一个hash中
	m := map[int]int{}
	for _, v := range small {
		m[v] = -1
	}

	for i := 0; i < bigLen; i++ {
		if _, ok := m[big[i]]; ok {
			if m[big[i]] == -1 {
				smallLen--
			}
			m[big[i]] = i
		}

		if smallLen <= 0 {
			minNum := getMinVal(m)
			if i-minNum+1 < result[1]-result[0]+1 {
				result[0] = minNum
				result[1] = i
			}
		}
	}
	// 不包含所有字符
	if smallLen > 0 {
		return []int{}
	}

	return result
}

func getMinVal(m map[int]int) int {
	minNum := math.MaxInt64
	for _, v := range m {
		minNum = min(minNum, v)
	}
	return minNum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// ==== 早上的 无重复字符的最长子串 ======
func lengthOfLongestSubstring(s string) int {
	var result int
	
	length := len(s)
	if length <= 0 {
		return result
	}

	// m:记录字符是否出现过,right:右指针
	m, right := map[byte]int{}, -1
	for i := 0; i < length; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		// m中没有出现
		for right+1 < length && m[s[right + 1]] == 0 {
			m[s[right+1]]++
			right++
		}

		result = max(result, right-i+1)
	}

	return result
}