package main

import (
	"fmt"
)

func main() {

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])

	// 마지막 2개 원소
	fmt.Println(aSlice[l-2 : l])
	fmt.Println(aSlice[l-2:])

	// 크기가 5인 용량
	t := aSlice[0:5:5]
	fmt.Println(t)

	// 추가하면 용량이 늘어남
	t = append(t, 1)
	fmt.Println(t)
}
