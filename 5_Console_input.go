package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What was the year you born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Printf("You will be %d yars old in 2020 ", 2020-input)

}
