/*
	88. 合并两个有序数组

	思路回顾
		定义两个指针: p1: nums1的尾部，p2:nums2的尾部; p = m+n-1

*/
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	p := m + n - 1
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
