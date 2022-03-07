package main

import "fmt"

func main() {
	numRows := 5
	fmt.Print(yanghuiTriangle(numRows))
}

func yanghuiTriangle(numRows int) map[int]map[int]int {
	arr := make(map[int]map[int]int)
	for i := 1; i <= numRows; i++ {
		arr[i] = make(map[int]int)
		arr[i][1] = 1
		for j := 2; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
		if i > 1 {
			arr[i][i] = 1
		}
	}
	return arr
}
