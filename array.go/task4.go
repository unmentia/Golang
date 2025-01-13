// Is n in the array?

package main
import "fmt"

func main(){
	nums := [] int {6, 25, 91, 23, 72, 9, 18, 6}
	n := 0
	fmt.Scan(&n)
	res := 0

	for i:=0; i<len(nums); i++ {
		if n == nums[i] {
			res++
		}
	}

	if res > 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}