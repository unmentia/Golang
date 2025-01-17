// The largest element

package main 

import "fmt"

func main(){
	nums := [] int {12, 21, 36, 71, 24, 85, 12, 30, 25} 
	max := nums[0]

	for i:=0; i<len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	fmt.Println(max)
}