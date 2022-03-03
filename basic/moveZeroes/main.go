package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 4, 0, 1, 2, 3, 0}
	fmt.Println(moveZeroes1(arr))
}

//移动所有0到末尾，为0的数据插入末尾
func moveZeroes(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i], arr[i+1:]...)
			arr = append(arr, 0)
		}
	}
	return arr
}

//移动所有0到末尾，不为0的数据依次前移
func moveZeroes1(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr
}
