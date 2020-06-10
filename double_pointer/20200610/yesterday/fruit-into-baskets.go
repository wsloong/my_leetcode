/*
904. 水果成篮

在一排树中，第 i 棵树产生 tree[i] 型的水果。
你可以从你选择的任何树开始，然后重复执行以下步骤：
	把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
	移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
	请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。
*/

/*
思路回顾
	维护一个移动滑块，每次迭代tree，就向滑块右移1位
	如果滑块的长度大于2个，就把滑块左侧-1，如果左侧滑块=0，则移除左侧滑块
*/
package main

func max(t1, t2 int) int {
	if t1 > t2 {
		return t1
	}
	return t2
}

// 定义一个结构用于存储值
type Counter struct {
	item map[int]int
}

func newCounter() *Counter {
	return &Counter{
		item: make(map[int]int),
	}
}

func (c *Counter) Get(k int) int {
	return c.item[k]
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

// 算法开始
func totalFruit(tree []int) int {
	// 定义滑块最左侧的值和结果
	i, result := 0, 0

	count := newCounter()

	// 迭代tree
	for j := range tree {
		count.Add(tree[j], 1) // 当前值加入到右侧滑块

		for count.Size() > 2 { // 最多只能有2个水果，当大于2个时候，将左侧值减一
			count.Add(tree[i], -1)
			if count.Get(tree[i]) == 0 { // 左侧滑块值没有了，删除掉
				count.Remove(tree[i])
			}
			i++ // 将滑块右移
		}

		result = max(result, j-i+1)
	}
	return result
}
