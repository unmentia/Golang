// The sum of numbers divisible by 5

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	res := 0

	for i:=0; i<=num; i++ {
		if i%5 == 0 {
			res+=i
		}
	}

	fmt.Println(res)
}