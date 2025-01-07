// The count of positive numbers between a and b

package main

import(
	"fmt"
)

func main(){
	num1 := 0
	num2 := 0
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	res := 0

	for i:=num1; i<=num2; i++ {
		if i>0 {
			res++
		}
	}

	fmt.Println(res)
}