// Sequence

package main

import(
	"fmt"
)

func main(){
	num1 := 0
	num2 := 0
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	mult := 0
	sum := 0

	for i:=0; i<num2; i++ {
		mult = mult*10+num1
		fmt.Println(mult)
		sum+=mult
	}

	fmt.Println(sum)
}