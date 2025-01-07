// The number of even numbers

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	res := 0

	for i:=0; i<num; i+=2 {
		res++
	}

	fmt.Println(res)
}