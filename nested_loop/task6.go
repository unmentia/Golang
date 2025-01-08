// List of numbers

package main

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	res := -1
	
	for y:=0; y<num; y++ {
		for x:=0; x<num; x++ {
			res++
			fmt.Printf("%d ", res)
		}
		fmt.Println()
	}
}