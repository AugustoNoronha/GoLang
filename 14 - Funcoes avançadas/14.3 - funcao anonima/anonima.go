package main

import "fmt"

func main() {
	func(num int) {
		fmt.Println("executado de func anoniicma")
		fmt.Println("Paramentro", 10)

	}(10)

	res := func() string {
		return "eu sou um retorno"
	}()
	fmt.Println("RES", res)
	//isso são as labdas ou arrow function

}
