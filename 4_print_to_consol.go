package main

import "fmt"

func main() {
	fmt.Printf("Hello %T %v %t", 10, 20, false)
	// store in variable
	var x string = fmt.Sprintf("Hello %T %v", 10, 20)
	fmt.Println(x)
	// test e f g t b x o
	fmt.Printf("Number: %g \n", 2.648538652354)
	fmt.Printf("Number: %e \n", 2.648538652354)
	fmt.Printf("Number: %f \n", 2.648538652354)
	fmt.Printf("Number: %x \n", 2.648538652354)
	fmt.Printf("Number: %t \n", 2.648538652354)
	fmt.Printf("Number: %b \n", 2.648538652354)
	fmt.Printf("Number: %o \n", 2.648538652354)
	fmt.Printf("Number: %s \n", "Shahriar ")
	fmt.Printf("Number: %q \n", "Shahriar ")
	// padding and percision
	fmt.Printf("Number: %5.2f \n", 52.525285)
	fmt.Printf("Number: %-9.2f is cool", 52.525285)
	fmt.Printf("Number: %07d", 45)

	
}
