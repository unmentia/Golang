// Reverse array

package main
import "fmt"

func main(){
	nums := [] int {1,2,3,4}
	arr := make([]int, len(nums)) 

	t := 0
	for i:=len(nums)-1; i>=0; i-- {
		arr[t] = nums[i]
		t++
	}

	fmt.Println(arr)
}