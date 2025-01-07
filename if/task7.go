// Positive and negative numbers

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	if num>0 {
		fmt.Println("Positive")
	} else if num<0 {
		fmt.Println("Negative")
	}
}