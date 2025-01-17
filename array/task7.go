// Add a number to the beginning

package main
import "fmt"

func main(){
	nums := [] int {0, 6, 25, 91, 23, 72, 9, 18, 6}
	n := 0
	fmt.Scan(&n)

	nums[0] = n
	fmt.Println(nums)
}