// A pair of numbers

package main

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	for y:=0; y<num; y++ {
		for x:=0; x<num; x++ {
			fmt.Printf("(%d,%d)", y, x)
		} 
		fmt.Println()
	}
}