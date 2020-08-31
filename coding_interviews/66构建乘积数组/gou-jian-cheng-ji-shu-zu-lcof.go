package main

func constructArr(a []int) []int {
	var res []int
	length := len(a)
	if length == 0 {
		return res
	}

	res = append(res, 1)
	tmp := 1

	for i := 1; i < length; i++ {
		res = append(res, res[i-1]*a[i-1])
	}

	for i := length - 2; i >= 0; i-- {
		tmp *= a[i+1]
		res[i] *= tmp
	}

	return res
}
