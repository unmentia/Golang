// a to the power of 5

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	res := 1

	for i:=0; i<5; i++ {
		res*=num
	} 
	fmt.Println(res)
}