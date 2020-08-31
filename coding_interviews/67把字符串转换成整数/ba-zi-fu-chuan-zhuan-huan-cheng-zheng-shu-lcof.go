package main

import (
	"fmt"
	"math"
)

func strToInt(str string) int {
	// res: 结果
	// sign: 符号
	res, i, sign, length, bndry := 0, 0, 1, len(str), math.MaxInt32/10

	if length == 0 {
		return 0
	}

	for str[i] == ' ' {
		i++
		if i == length {
			return 0
		}
	}

	// 开头是 '-' 标示是负数
	if str[i] == '-' {
		sign = -1
	}

	if str[i] == '-' || str[i] == '+' {
		i++
	}

	for j := i; j < length; j++ {
		if str[j] < '0' || str[j] > '9' {
			break
		}

		if res > bndry || res == bndry && str[j] > '7' {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		res = res*10 + int(str[j]-'0')
	}
	return sign * res

}

func main() {
	str := " "
	fmt.Println(strToInt(str))
}
