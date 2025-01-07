//  A number that is divisible by 3 and 4

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	if num%3 == 0 && num%4 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}