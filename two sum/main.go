package main

import (
	"fmt"
)

func find_pair_sum1(nums []int, target int) (int, int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return nums[i], nums[j]
				//fmt.Println(nums[i],nums[j])
			}
		}

	}

	return -1, -1
}

func main() {
	nums := []int{4, 3, 2, 5, 1, 10, 7, 100, 23, 27}
	target := 9
	p1,p2 := find_pair_sum1(nums, target)
	fmt.Println(p1, p2)
}
