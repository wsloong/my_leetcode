package main

import (
	"fmt"
	"strings"
)

func minNumber(nums []int) string {
	var strs []string
	for _, num := range nums {
		strs = append(strs, fmt.Sprintf("%d", num))
	}

	var fastSort func(left, right int)
	fastSort = func(left, right int) {
		if left >= right {
			return
		}

		i, j := left, right

		for i < j {
			for (strs[j]+strs[left] >= strs[left]+strs[j]) && i < j {
				j--
			}

			for (strs[i]+strs[left] <= strs[left]+strs[i]) && i < j {
				i++
			}
			strs[i], strs[j] = strs[j], strs[i]
		}
		strs[i], strs[left] = strs[left], strs[i]
		fastSort(left, i-1)
		fastSort(i+1, right)
	}

	fastSort(0, len(strs)-1)
	return strings.Join(strs, "")
}
