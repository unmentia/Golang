// A large number

package main

import(
	"fmt"
)

func main(){
	num1 := 0
	num2 := 0
	num3 := 0
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)

	if num1>num2 && num1>num3 {
		fmt.Println(num1)
 	} else if num2>num1 && num2>num3 {
		fmt.Println(num2)
	} else if num3>num1 && num3>num2 {
		fmt.Println(num3)
	}
}