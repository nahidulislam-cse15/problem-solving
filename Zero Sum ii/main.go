package main

import (
	"fmt"
)
type startend struct {
	start int
	end int
}

// 0(n^3 ) naive approach as there are n2 subarrays in an array of size n, and it takes O(n) time to find the sum of its elements
func hasZeroSumSubarray1(nums []int, n int) bool {
	var sum, flag int
	//mp := make(map[int]int)
	var subarray []startend
	k,cnt:=0,0
	for i := 0; i < len(nums); i++ {
		sum = 0

		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == 0 {
				var se startend
				fmt.Println("Sum found between indexes :", i, "and ", j)
				se.start=i
				se.end=j
				subarray=append(subarray,se)
				k++
				cnt++
				// fmt.Print("{")
				// for k:=i ; k <= j; k++{
				// 	fmt.Printf("%d ",nums[k] ) //printing subarray value //
				// }
				// fmt.Println("}")
				//endindex=append(endindex,j)
				flag = 1
			}

		}
	}
	// for key, value := range mp {
	// 	for i := key; i <= value; i++ {
	// 		fmt.Printf("%d ", nums[i])
	// 	}
	// 	
	// }
	for i :=0; i< cnt; i++ {
	for j:= subarray[i].start; j<= subarray[i].end; j++ {
		fmt.Printf("%d ",nums[j])
	}
	fmt.Println()
}
	if flag == 1 {
		return true
	} else {
		return false

	}
}

//hasing  o(n)->T.c -S.c
func hasZeroSumSubarray2(nums []int, n int) bool {
	mp := make(map[int]bool)
	sum := 0
	mp[0] = true
	for i := 0; i < len(nums)-1; i++ {
		sum += nums[i]
		if mp[sum] {
			return true
		} else {
			mp[sum] = true
		}
	}
	return false

}

func main() {
	// nums:=[] int{ -3, 2, 3, 1, 6  };
	//nums := []int{4, 2, -3, -1, 0, 4}
	nums := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}

	fmt.Println(hasZeroSumSubarray1(nums, len(nums)))
	//fmt.Println(hasZeroSumSubarray2(nums, len(nums)))

}
