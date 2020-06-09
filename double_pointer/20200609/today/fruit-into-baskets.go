/*
904. 水果成篮

在一排树中，第 i 棵树产生 tree[i] 型的水果。
你可以从你选择的任何树开始，然后重复执行以下步骤：
	把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
	移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
	请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。

你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
用这个程序你能收集的水果总量是多少？

示例 1：
	输入：[1,2,1]
	输出：3
	解释：我们可以收集 [1,2,1]。

示例 2：
	输入：[0,1,2,2]
	输出：3
	解释：我们可以收集 [1,2,2].
	如果我们从第一棵树开始，我们将只能收集到 [0, 1]。

示例 3：
	输入：[1,2,3,2,2]
	输出：4
	解释：我们可以收集 [2,3,2,2].
	如果我们从第一棵树开始，我们将只能收集到 [1, 2]。

示例 4：
	输入：[3,3,3,1,2,1,1,2,3,3,4]
	输出：5
	解释：我们可以收集 [1,2,1,1,2].
	如果我们从第一棵树或第八棵树开始，我们将只能收集到 4 个水果。
 
提示：
1 <= tree.length <= 40000
0 <= tree[i] < tree.length

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fruit-into-baskets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func totalFruit(tree []int) int {
	result, i := 0, 0

	count := newCounter()

	for j := range tree {
		count.Add(tree[j], 1)
		for count.Size() >= 3 {
			count.Add(tree[i], -1)
			if count.Get(tree[i]) == 0 {
				count.Remove(tree[i])
			}
			i++
		}
		result = max(result, j - i + 1)
	}

	return result
}

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}

	return v2
}

type Counter struct {
	item map[int]int
}

func newCounter() *Counter {
	return &Counter{
		item: make(map[int]int),
	}
}

func (c *Counter) Get(k int) int {
	value, ok := c.item[k]
	if ok {
		return value
	}
	return 0
}

func (c *Counter) Add(k, v int) {
	c.item[k] = c.Get(k) + v
}

func (c *Counter) Size() int {
	return len(c.item)
}

func (c *Counter) Remove(k int) {
	delete(c.item, k)
}


// ====== 早上的  删除排序数组中的重复项 ===
/*
思路回顾
	题目虽然是删除但是跟确切来说是交换
	使用双指针，i, j 
	如果nums[i] != nums[j], i++, nums[i] = nums[j]
*/

func removeDuplicates(nums []int) int {
	length, i := len(nums), 0
	for j := 1; j < length; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i]= nums[j]
		}
	}
	return i + 1
}