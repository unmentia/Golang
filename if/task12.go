// Lowercase or uppercase alphabet

package main

import(
	"fmt"
)

func main(){
	var inp string
	fmt.Scan(&inp)

	if inp>="a" && inp<="z" {
		fmt.Println("Lowercase")
	} else if inp>="A" && inp<="Z" {
		fmt.Println("Uppercase")
	} else {
		fmt.Println("None")
	}
}