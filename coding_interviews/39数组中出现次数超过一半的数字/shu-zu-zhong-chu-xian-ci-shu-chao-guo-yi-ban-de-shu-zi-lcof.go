package main

/*
哈希表统计法： 遍历数组 `nums`，用 `HashMap` 统计各数字的数量，最终超过数组长度一半的数字则为众数。此方法时间和空间复杂度均为 O(N)。
数组排序法： 将数组 `nums` 排序，由于众数的数量超过数组长度一半，因此 `数组中点的元素` 一定为众数。此方法时间复杂度 O(NlogN)。
摩尔投票法： 核心理念为 “正负抵消” ；时间和空间复杂度分别为 `O(N) 和 `O(1)` ；是本题的最佳解法。

摩尔投票法：
	由于众数出现的次数超过数组长度的一半，若记`众数`的票数为`+1`，`非众数`的票数为`-1`.则一定由于所有数字的 `票数和 > 0`

作者：jyd
链接：https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/solution/mian-shi-ti-39-shu-zu-zhong-chu-xian-ci-shu-chao-3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func majorityElement(nums []int) int {
	// 初始化票数为0
	var votes int
	// 初始化众数，默认为第一个数字，随着循环改变
	var x int

	for _, num := range nums {
		// 当投票为0(初始化或者相互抵消), 众数等于当前数字
		if votes == 0 {
			x = num
		}

		// 如果当前数等于众数，投票+1，否则投票-1
		if num == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}

/*
拓展：
1.由于题目明确给定 `给定的数组总是存在多数元素`，因此本题不用考虑 `数组中不存在众数` 的情况。
2. 若考虑，则需要加入一个 “验证环节” ，遍历数组 nums 统计 x 的数量。
	1.若 x 的数量超过数组长度一半，则返回 x ；
	2.否则，返回 0（这里根据不同题目的要求而定）。
*/

func majorityElementWithCheck(nums []int) int {
	votes, count, x := 0, 0, 0

	for _, num := range nums {
		if votes == 0 {
			x = num
		}

		if num == x {
			votes++
		} else {
			votes--
		}
	}

	// 验证
	for _, num := range nums {
		if num == x {
			count++
		}
	}

	if count > len(nums)/2 {
		return x
	}

	// 没有众数时候返回0
	return 0
}
