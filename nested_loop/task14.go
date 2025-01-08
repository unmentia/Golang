// Parallelogram

package main 

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	for y:=0; y<num; y++ {
		for x:=0; x<num+y; x++ {
			if y==x || y<x {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}