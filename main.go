package main

import (
	"fmt"

	"dont-go/functions"
)

func main() {

	x := 4
	x = functions.Increment(x)
	fmt.Println(x)
	fmt.Println(functions.Concat("Abcd", "Efgh"))

	firstName, _ := functions.ReturnName("Prani", "Daemon")

	fmt.Println(firstName)
}
