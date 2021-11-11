package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Name: ")
	scanner.Scan()
	var name = scanner.Text()

	fmt.Print("Age: ")
	scanner.Scan()
	var age, _ = strconv.Atoi(scanner.Text())

	var count = age * 365 // counting how many days the user live

	fmt.Println("Your name:", name)
	fmt.Println("Your age:", age)
	fmt.Println("Age in days:", count, "days")

}
