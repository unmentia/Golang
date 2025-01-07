// A large number

package main

import(
	"fmt"
)

func main(){
	num1 := 0
	num2 := 0
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	if num1>num2 {
		fmt.Println(num1)
	} else {
		fmt.Println(num2)
	}
}