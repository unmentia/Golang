// Diagonal

package main

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	for y:=0; y<num; y++ {
		for x:=0; x<y+1; x++ {
			if x==y {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}