// a to the power of n 

package main

import(
	"fmt"
)

func main(){
	num := 0
	pw := 0
	fmt.Scan(&num)
	fmt.Scan(&pw)

	res := 1

	for i:=0; i<pw; i++ {
		res*=num
	}
	fmt.Println(res)
}