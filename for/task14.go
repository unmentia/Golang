// A perfect number

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	res := 0

	for i:=1; i<num; i++ {
		if num%i == 0 {
			res+=i
		}
	}
	
	if res==num {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}