// Columns in pairs of numbers

package main

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	for y:=0; y<num; y++ {
		for x:=0; x<num; x++ {
			if y==0 || y==num-1 || x==0 || x%2==0 {
				fmt.Printf("(%d,%d)", y, x)
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}