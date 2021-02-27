package main

import "fmt"

func main() {
	age  := 13
	fmt.Println("you are: ", age)
	if age >= 18 {
		fmt.Println("you can ride alone ")
	}else{
		fmt.Println("Sorry you can not ride it ")
		fmt.Printf("Wait %d years \n", 18-age)
	}
	
	fmt.Println("enjoy the Day ! ")
}
