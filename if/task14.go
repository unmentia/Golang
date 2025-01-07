// Interior angles of a triangle

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

	if num1>0 && num2>0 && num3>0 && num1+num2+num3 == 180 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}