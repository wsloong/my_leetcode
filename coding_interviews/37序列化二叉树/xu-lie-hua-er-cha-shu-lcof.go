package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	// 定义辅助队列
	queue := []*TreeNode{root}
	var result []string

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if root == nil {
			result = append(result, "null")
		} else {
			result = append(result, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left, node.Right)
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(result, ","))

}

func deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// values: 节点的值
	// i: 对应当前节点的索引
	values, i := strings.Split(data[1:len(data)-1], ","), 1

	// 根节点
	rootVal, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: rootVal}

	// 定义赋值的队列
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 左节点
		if values[i] != "null" {
			leftVal, _ := strconv.Atoi(values[i])
			node.Left = &TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		i++

		// 右节点
		if values[i] != "null" {
			rightVal, _ := strconv.Atoi(values[i])
			node.Right = &TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
