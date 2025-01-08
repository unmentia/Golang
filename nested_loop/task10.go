// Right-angled triangle

package main 

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	for y:=0; y<num; y++ {
		for x:=0; x<=y; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}