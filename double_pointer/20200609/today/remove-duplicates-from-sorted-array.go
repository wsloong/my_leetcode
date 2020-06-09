/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1:
	给定数组 nums = [1,1,2],
	函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
	你不需要考虑数组中超出新长度后面的元素。
示例 2:
	给定 nums = [0,0,1,1,1,2,2,3,3,4],
	函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
	你不需要考虑数组中超出新长度后面的元素。

说明:
	为什么返回数值是整数，但输出的答案是数组呢?
	请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
	你可以想象内部操作如下:

	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	int len = removeDuplicates(nums);

	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
	for (int i = 0; i < len; i++) {
	    print(nums[i]);
	}

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	// 定义两个指针，慢指针i:0，快指针j:1
	// 如果 i的值和j的值相等，增加j，跳过重复项
	// 如果不等，就让i的值+1，nums[i] = nums[j]
	var i int
	for j := 1; j < length; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 该解答错误，没有仔细审题，
// 原题说 `原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4`
// 就是直接修改原数组的值，让前x个数值不同，并返回x
func removeDuplicatesERROR(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	// 新数组的长度，默认为数组长度-1，随着数组相同元素的删除会减少
	currentLength := length

	for i := 0; i < currentLength; i++ {
		// 如果当前值和下一个元素相同
		// 不停的后移j知道不和i相等
		j := i
		for j < currentLength && nums[i] == nums[j] {
			j++
		}

		if j != i {
			nums = append(nums[:i+1], nums[j:currentLength]...)
			currentLength = currentLength - (j - 1 - i)
		}
	}

	return currentLength
}
