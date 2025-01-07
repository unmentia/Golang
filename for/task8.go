// The sum and the count of numbers divisible by 3 between the numbers a and b

package main

import(
	"fmt"
)

func main(){
	num1 := 0
	num2 := 0
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	sum := 0
	count := 0

	for i:=num1; i<num2; i++ {
		if i%3 == 0 {
			sum+=i
			count++
		}
	}

	fmt.Println("Sum =", sum)
	fmt.Println("Count =", count)
}