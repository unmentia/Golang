// Sequence

package main

import "fmt"

func main(){
	num := 0
	fmt.Scan(&num)

	var res float32

	for i:=1; i<=num; i++ {
		res += 1/float32(i)
	}

	fmt.Printf("%f\n", res)
}	