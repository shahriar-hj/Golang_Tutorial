package main

import "fmt"

func main() {

	var number = 46.54
	number = number + 45
	fmt.Printf("%T", number)
	// another way to sset variable
	number_2 := 34
	fmt.Printf("%T", number_2)

	not_number := "Hello Shahriar"
	fmt.Printf("%T", not_number)
}

// print format of the Variable
