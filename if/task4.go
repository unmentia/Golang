// Even or odd number

package main

import (
	"fmt"
)

func main(){
	num := 0;
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}