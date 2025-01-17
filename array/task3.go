// Average value

package main
import "fmt"

func main(){
	nums := [] int {6, 25, 91, 23, 72, 9, 18, 6}
	res := 0

	for i:=0; i<len(nums); i++ {
		res += nums[i]
	}

	fmt.Println(res/len(nums))
}