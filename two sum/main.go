package main

import (
	"fmt"
	"sort"
)

//brute force
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

// binary search
// binaryserach

func find_pair_sum2(nums []int, target int) (int, int) {
	//sort.Sort(nums)
	sort.Ints(nums)
	low, high := 0, len(nums)-1
	for low < high {

		if nums[low]+nums[high] == target {
			return nums[low], nums[high]
		} else if nums[low]+nums[high] < target {
			low++
		} else {

			high--

		}
	}
	return -1, -1
}

//map hashing
func find_pair_sum3(nums []int, target int) (int, int) {
	mp := make(map[int]int)//empty map
	//if(mp[nums] )

	for i := 0; i < len(nums)-1; i++ {
		if mp[nums[i]] == target-nums[i] {
			return mp[nums[i]], nums[i]
		} else {
			mp[target-nums[i]] = nums[i]
		}

	}

	return -1, -1
}

func main() {
	nums := []int{4, 3, 2, 5, 1, 10, 7, 100, 23, 27}
	target := 27

	//p1,p2 := find_pair_sum1(nums, target)
	//p1, p2 := find_pair_sum2(nums, target)
	p1, p2 := find_pair_sum3(nums, target)
	fmt.Println(p1, p2)
}
