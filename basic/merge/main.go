package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	fmt.Print(merge(nums1, nums2, m, n))
}

func merge(nums1 []int, nums2 []int, m int, n int) []int {
	nums1 = append(append(nums1[:m]), append(nums2[:n])...)
	sort.Ints(nums1)
	return nums1
}
