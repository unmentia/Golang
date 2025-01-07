// Alpha or number

package main

import(
	"fmt"
)

func main(){
	var inp string
	fmt.Scan(&inp)

	if inp>="a" && inp<="z" || inp>="A" && inp<="Z" {
		fmt.Println("Alpha")
	} else {
		fmt.Println("Number")
	}
}