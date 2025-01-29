// Add b to index a

package main
import "fmt"

func main(){
	nums := [] int {4, 5, 6, 7}
	arr := make([]int, len(nums)+1)

	var a, b int
	fmt.Scan(&a, &b)

	for i:=0; i<len(arr); i++ {
		if i<a {
			arr[i] = nums[i]
		} else if i==a {
			arr[i] = b
 		} else {
			arr[i] = nums[i-1]
		}
	}

	fmt.Println(arr)
}