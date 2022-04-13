package main
import (
	"fmt"
)
// 0(n^2) naive approach
func hasZeroSumSubarray1(nums []int , n int) bool {
	var sum int
	for i := 0; i < len(nums)-1; i++{
		sum=nums[i]
	for j := i+1; j < len(nums); j++{
		if(sum==0){
			return true
		}
		
	  sum +=  + nums[j];


	}
}
return false
  
  } 
  //hasing  o(n)->T.c -S.c
  func hasZeroSumSubarray2(nums []int , n int) bool {
  mp:=make(map[int]bool)
  sum:=0
  mp[0]=true
  for i := 0; i < len(nums)-1; i++ {
	  sum+=nums[i]
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
	 nums:=[] int{  4, 2, -3, -1, 0, 4 };
	// nums:=[] int{ 3, 4, -7, 3, 1, 3, 1, -4, -2, -2 };
	
    fmt.Println(hasZeroSumSubarray1(nums, len(nums)))
	fmt.Println(hasZeroSumSubarray2(nums, len(nums)))

}