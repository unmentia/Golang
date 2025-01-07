// Leap year

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	if num%4 == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}