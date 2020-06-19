package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	// return strings.ReplaceAll(s, " ", "%20")
	var result []rune

	for _, v := range s {
		if v == ' ' {
			result = append(result, []rune("%20")...)
		} else {
			result = append(result, v)
		}
	}

	return string(result)
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
