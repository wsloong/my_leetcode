/* 面试题 17.11. 单词距离
有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。
如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

示例：
	输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
	输出：1
提示：
	words.length <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-closest-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func findClosest(words []string, word1 string, word2 string) int {
	t1, t2, result := -1, -1, len(words)

	for i := range words {
		if words[i] == word1 {
			t1 = i
		} else if words[i] == word2 {
			t2 = i
		}

		if t1 != -1 && t2 != -1 {
			result = min(result, abs(t1, t2))
		}

		if result == 1 {
			return result
		}
	}

	return result
}

func min(t1, t2 int) int {
	if t1 < t2 {
		return t1
	}
	return t2
}

func max(t1, t2 int) int {
	if t1 > t2 {
		return t1
	}
	return t2
}

func abs(t1, t2 int) int {
	return max(t1, t2) - min(t1, t2)
}
