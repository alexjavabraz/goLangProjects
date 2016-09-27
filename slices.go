package main

import (
	"fmt"
)

func main() {

	fmt.Println("imprime")

	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 22
	fmt.Println(slice)

	s1 := slice[:3]
	fmt.Println(s1)
}
