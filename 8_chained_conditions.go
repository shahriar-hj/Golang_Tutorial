package main

import "fmt"

// || ! &&
func main() {

	val := 5<7 || 7 > 9
	fmt.Printf("%t \n", val)
	val2 := 6>7 && 3<1
	fmt.Printf("%t \n", val2)
	
}
