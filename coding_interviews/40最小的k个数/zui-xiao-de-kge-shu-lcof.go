package main

import (
	"fmt"
)

func getLeastNumbers(arr []int, k int) []int {
	var result []int
	if k == 0 || len(arr) == 0 {
		return result
	}

	// 参数 `k-1`：要找的是`索引为k-1的数`
	quickSearch(arr, 0, len(arr)-1, k-1)

	result = make([]int, k)
	copy(result, arr)
	return result
}

func quickSearch(arr []int, low, hi, k int) {
	if low == hi {
		return
	}

	i, j, pivot := low, hi, arr[low]

	for i <= j {
		for i <= j && arr[i] < pivot {
			i++
		}

		for i <= j && arr[j] > pivot {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	if low <= k && k <= j {
		quickSearch(arr, low, j, k)
	} else {
		quickSearch(arr, i, hi, k)
	}
}

func main() {
	a := []int{3, 2, 1}
	k := 2
	fmt.Println(getLeastNumbers(a, k))
}
