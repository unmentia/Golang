// n-th index

package main
import "fmt"

func main(){
	nums := [] int {6, 25, 91, 23, 72, 9, 18, 6}
	inx := 0
	fmt.Scan(&inx)
	res := nums[inx]

	for i:=0; i<len(nums); i++ {
		if inx == i {
			res = nums[inx]
		}
	}

	fmt.Println(res) 
}