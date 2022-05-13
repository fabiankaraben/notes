package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 1, 2, 4}

	fmt.Println(arr)
	arr = reverse(arr)
	fmt.Println(arr)

}

func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		opI := len(s) - 1 - i
		s[i], s[opI] = s[opI], s[i]
	}

	return s
}
