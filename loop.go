package main

import "fmt"

func main() {

	number := 20

	for number > 0 {
		fmt.Println(number)
		number--
	}

	//LIKE WHILE
	fmt.Println("loop while")
	for number := 2; number >= 0; number-- {
		fmt.Println(number)
	}

}
