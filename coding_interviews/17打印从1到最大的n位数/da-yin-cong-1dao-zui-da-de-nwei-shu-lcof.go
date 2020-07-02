package main

func printNumbers(n int) []int {
	var result []int

	var end int
	for n > 0 {
		end = 10*end + 9
		n--
	}

	for i := 1; i <= end; i++ {
		result = append(result, i)
	}
	return result
}
