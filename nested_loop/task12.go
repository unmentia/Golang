// Triangle of numbers

package main

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	for y:=1; y<=num; y++ {
		for x:=1; x<=y; x++ {
			fmt.Printf("%d", x)
		}
		fmt.Println()
	}
}