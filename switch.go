package main

import (
	"fmt"
)

func main() {

	number := 5

	if n:= 5; n > 0 {
		fmt.Printf("MAIOR QUE ZERO")
	}else{
		fmt.Printf("MENOR QUE ZERO")
	}

	fmt.Println("")
	fmt.Println("switch")

	switch number {
	case 1:
		fmt.Println("vale 1")
	case 2:
		fmt.Println("vale 2")
	default:
		fmt.Println("default")
	}
}
