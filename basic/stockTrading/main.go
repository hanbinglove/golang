package main

import "fmt"

func main() {
	arr := []int{2, 5, 2, 1, 2, 6, 457, 2, 2345, 2345, 5, 16, 67, 23445, 1, 2345, 234}
	fmt.Print(stockTrading(arr))
}

func stockTrading(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	max := 0
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			max = 0
		} else {
			max = mathMax(max, arr[i]-min)
		}
	}
	return max
}

func mathMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
