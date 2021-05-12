package main

import "fmt"

func reverseArray(a []int32) []int32 {
	result := []int32{}
	for i := 1; i <= len(a); i++ {
		result = append(result, a[len(a)-i])
	}
	return result
}

func main() {
	testArr := []int32{5, 4, 3, 2, 1}
	fmt.Println(reverseArray(testArr))
}
