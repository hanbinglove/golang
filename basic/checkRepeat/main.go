package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 0, 4, 0, 1, 2, 3, 0}
	fmt.Println(checkRepeat(arr))
	fmt.Println(checkRepeat1(arr))
}

//先进行数组排序，再进行依次对比
func checkRepeat(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}

//使用哈希表记录元素是否出现过
func checkRepeat1(arr []int) bool {
	newArray := map[int]bool{}
	for _, item := range arr {
		if _, ok := newArray[item]; ok {
			return true
		}
		newArray[item] = true
	}
	return false
}
