// Add a number to the end

package main
import "fmt"

func main(){
	nums := [] int {6, 25, 91, 23, 72, 9, 18, 6, 0}
	n := 0
	fmt.Scan(&n)

	nums[len(nums)-1] = n

	fmt.Println(nums)
}