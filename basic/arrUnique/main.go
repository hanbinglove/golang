package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 0, 4, 0, 1, 2, 3, 0}
	fmt.Println(arrUnique1(arr))
}

//数组排重-先排序，再依次对比删除
func arrUnique(arr []int) []int {
	sort.Ints(arr)
	j := 0
	for {
		if arr[j] == arr[j+1] {
			arr = append(arr[:j+1], arr[j+2:]...)
		} else {
			j++
		}
		if j >= len(arr)-1 {
			break
		}
	}
	return arr
}

//赋值给一个新的数组
func arrUnique1(arr []int) []int {
	newArr := map[int]bool{}
	i := 0
	for {
		if _, ok := newArr[arr[i]]; ok {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			newArr[arr[i]] = true
			i++
		}
		if i >= len(arr) {
			break
		}
	}
	return arr
}
