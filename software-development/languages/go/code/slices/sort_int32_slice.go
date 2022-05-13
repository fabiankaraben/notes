package main

import (
	"fmt"
	"sort"
)

type int32Slice []int32

func (arr int32Slice) Len() int {
	return len(arr)
}

func (arr int32Slice) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr int32Slice) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := int32Slice{3, 1, 2, 4}

	fmt.Println(arr)
	sort.Sort(arr)
	fmt.Println(arr)
}
