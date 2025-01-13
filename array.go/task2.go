// The smallest element

package main
import "fmt"

func main(){
	nums := [] int {12, 21, 36, 71, 24, 85, 1, 30, 25}
	min := nums[0]

	for i:=0; i<len(nums); i++ {
		if min>nums[i] {
			min = nums[i]
		}
	}

	fmt.Println(min)
}