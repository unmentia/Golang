// Day of the week

package main

import(
	"fmt"
)

func main(){
	num := 0
	fmt.Scan(&num)

	if num==1 {
		fmt.Println("Monday")
	} else if num==2 {
		fmt.Println("Tuesday")
	} else if num==3 {
		fmt.Println("Wednesday")
	} else if num==4 {
		fmt.Println("Thursday")
	} else if num==5 {
		fmt.Println("Friday")
	} else if num==6 {
		fmt.Println("Saturday")
	} else if num==7 {
		fmt.Println("Sunday")
	}
}