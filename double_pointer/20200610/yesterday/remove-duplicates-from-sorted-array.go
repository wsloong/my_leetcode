/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/
/*
思路回顾
定义两个指针i(从0开始),j(从1开始)
如果nums[i] != nums[j], i++, nums[i] = nums[j]
如果相等，就继续循环
*/

func removeDuplicates(nums []int) int {
	i, length := 0, len(nums)

	for j := 1; j < length; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i+1
}