// A number that is divisible by 5

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	if num%5 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}