package main

import (
	"fmt"
)

func main() {
	a := 11
	if a%2 == 0 {
		fmt.Println(a, "is even!")
	} else {
		fmt.Println(a, "is odd!")
	}

	b := 100
	fmt.Printf("B is %d", b)

}
