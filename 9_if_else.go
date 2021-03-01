package main

import "fmt"

func main() {
	age := 19
	fmt.Println("you are: ", age)
	if age >= 18 {
		fmt.Println("you can ride alone! ")
	} else if age >= 14 {
		fmt.Println("you can ride with parent! ")

	} else {
		fmt.Println("Sorry, you can not Ride!")
		fmt.Printf("Wait %d years \n", 18-age)

	}

	fmt.Println("enjoy the Day ! ")
}
