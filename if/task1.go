// Comparison with 10

package main

import (
	"fmt"
)

func main () {
	a := 0
	fmt.Scan(&a)

	if a>10 {
		fmt.Println(a+3)
	} else if a<10 {
		fmt.Println(a*2) 
	} else {
		fmt.Println(22)
	}
}