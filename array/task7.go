// Add a number to the beginning

package main
import "fmt"

func main(){
	nums := [] int {0, 6, 25, 91, 23, 72, 9, 18, 6}
	arr := make([]int, len(nums)+1)
	n := 0
	fmt.Scan(&n)

	arr[0] = n
	for i:=0; i<len(nums); i++ {
		arr[i+1] = nums[i]
	}
	fmt.Println(arr)
}