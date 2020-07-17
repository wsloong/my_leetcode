package main

func validateStackSequences(pushed []int, popped []int) bool {
	// 辅助栈： 用于保存pushed的元素
	// 当刚压入的元素 = popped[i]时候，将该元素弹出
	stack, i := []int{}, 0

	for _, num := range pushed {
		stack = append(stack, num)
		// 当辅助栈元素大于0，并且栈顶的元素 = popped[i]的值
		// 弹出 stack 顶部的元素
		// i++
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	// 如果最后 stack为空，返回 true
	return len(stack) == 0
}
