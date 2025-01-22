// Add a number to the end

package main
import "fmt"

func main(){
	nums := [] int {6, 25, 91, 23, 72, 9, 18, 6}
	arr := make([]int, len(nums)+1)
	n := 0
	fmt.Scan(&n)

	for i:=0; i<len(arr); i++ {
		if i<len(arr)-1 {
			arr[i] = nums[i]
		} else {
			arr[i] = n
		}
	}
	fmt.Println(arr)
}