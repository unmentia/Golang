// The sum of odd numbers

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	res := 0

	for i:=1; i<=num; i+=2 {
		res+=i
	}

	fmt.Println(res)
}