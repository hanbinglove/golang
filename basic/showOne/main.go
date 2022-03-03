package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 0, 4, 0, 1, 2, 3, 0}
	fmt.Println(showOne1(arr))
}

//通过计数器进行统计出现的次数再进行过滤
func showOne(arr []int) (resArr []int) {
	newArr := map[int]int{}
	for _, item := range arr {
		if _, ok := newArr[item]; ok {
			newArr[item] = 2
		} else {
			newArr[item] = 1
		}
	}
	for item, num := range newArr {
		if num == 2 {
			continue
		}
		resArr = append(resArr, item)
	}
	return resArr
}

//先排序，再挨个数字找
func showOne1(arr []int) []int {
	sort.Ints(arr)
	i := 0
	for {
		j := 0
		for {
			if arr[i] == arr[i+1] {
				arr = append(arr[:i+1], arr[i+2:]...)
				j++
			} else {
				break
			}
		}
		if j > 0 {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
		if i >= len(arr)-1 {
			break
		}
	}
	return arr
}
