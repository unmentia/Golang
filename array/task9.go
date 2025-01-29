// Remove the n-th index

package main
import "fmt"

func main(){
	nums := [] int {4,5,6,7}
	arr := make([]int, len(nums)-1)

	var a int 
	fmt.Scan(&a)

	t := 0
	for i:=0; i<len(nums); i++ {
		if i==a {
			continue 
		} else {
			arr[t] = nums[i]
			t++
		}
	}

	fmt.Println(arr)
}