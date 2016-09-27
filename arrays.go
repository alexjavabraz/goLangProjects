package main

import (
	"fmt"
)

func main() {
	taboada := [10]int{}

	fmt.Println("taboada ", taboada)

	taboada = [10]int{0, 5, 10}

	fmt.Println("taboada ", taboada)

	taboada[3] = 15

	fmt.Println("taboada ", taboada)

	fmt.Println("tamanho do array ", len(taboada))

	userold := map[string]string{}

	user := map[string]string{
		"name":   "alex",
		"idade":  "34",
		"altura": "180",
	}

	fmt.Println(userold)

	fmt.Println(user)

	user["altura"] = "1,80"
	fmt.Println(user)

	fmt.Println(len(user))

	fmt.Println(len(user["name"]))

	user["teste"] = "teste"

	fmt.Println(user)

	slice := []int{0, 5, 10}
	slice[2] = 22

	slice = append(slice, 90, 110, 50)

	fmt.Println(slice)

}
