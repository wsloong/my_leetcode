package main

import (
	"fmt"
	"sort"
)

// 核心点
// 除去大小王，max - min < 5
func isStraight(nums []int) bool {
	sort.Ints(nums)
	var joker int

	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}

	return nums[4]-nums[joker] < 5 // 最大值 - 最小值 < 5; 则可以构成顺子
}

func main() {

	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println(isStraight(nums1))

	nums2 := []int{0, 0, 1, 2, 5}
	fmt.Println(isStraight(nums2))
}
