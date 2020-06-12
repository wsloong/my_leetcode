/*
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明:
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
输入:
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

// 最优解
// 官方提供了一个思路，是从后往前，进行扫描。这样不需要重新申请m个空间
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 申请两个指针，分别从nums1，nums2的最后开始
	p1, p2 := m-1, n-1
	p := m + n - 1 // 总长度，nums1的所有长度，包括额外空间

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}

	copy(nums1[:p2+1], nums2[:p2+1])
}

/*
将nums1的前m个数值copy出来，放到一个新的数组中
定义p, q 两个指针，分别对应nums1和nums2的开头部分
移动双指针进行一次移动进行比较

// 由于go的数组的特殊性，不能使用 nums1Copy = nums1[:m]
// nums1 不能重置为 []int{}，需要使用下标一一替换，因为这里是在元素组上进行修改，没有返回值
*/

func mergeFromBefore(nums1 []int, m int, nums2 []int, n int) {
	var nums1Copy []int
	copy(nums1Copy, nums1[:m])

	p, q := 0, 0
	var index int

	for p < m && q < n {
		if nums1Copy[p] < nums2[q] {
			nums1[index] = nums1Copy[p]
			p++
		} else {
			nums1[index] = nums2[q]
			q++
		}
		index++
	}

	if p < m {
		nums1[index] = nums1Copy[p]
		p++
		index++
	}

	for q < n {
		nums1[index] = nums2[q]
		q++
		index++
	}
}
