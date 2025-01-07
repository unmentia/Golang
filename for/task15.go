// The square of a sequence

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	sum := 0

	for i:=1; i<=num; i++ {
		sum+=i*i
	}

	fmt.Println(sum)
}