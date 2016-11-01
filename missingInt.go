package main

import "fmt"

func main() {
	var ans int
	arr := [10]int{4, 1, 0, 2, 9, 6, 8, 7, 5, 3}
	otherArr := [9]int{4, 6, 7, 2, 1, 0, 8, 3, 9}

	ans = missingInt(arr[:], otherArr[:])
	fmt.Println(ans)
}

func missingInt(arr1 []int, arr2 []int) int {
	sum1 := 0
	sum2 := 0

	for _, value := range arr1 {
		sum1 += value
	}

	for _, value := range arr2 {
		sum2 += value
	}

	return sum1 - sum2
}