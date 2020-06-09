/*
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。
示例 1:

输入: S = "ababcbacadefegdehijhklij"
输出: [9,7,8]
解释:
	划分结果为 "ababcbaca", "defegde", "hijhklij"。
	每个字母最多出现在一个片段中。
	像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

注意:
	S的长度在[1, 500]之间。
	S只包含小写字母'a'到'z'
*/

/*
思路回顾
1 遍历字符串，记录每个字符最后出现的位置
2 使用start, end，标示要划分的区间，根据最后出现的位置决定 end的值
3 当遍历到i = end 时候，就把值放到结果中。然后将 start = end+1,找下一个区间
*/
package main

func partitionLabels(S string) []int {
	last := make(map[rune]int)

	for i, s := range S {
		last[s] = i
	}

	var start, end int // 记录区间
	var result []int

	for i, s := range S {
		if last[s] > end {
			end = last[s]
		}

		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
