package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan Nama: ")
	scanner.Scan()
	var name = scanner.Text()

	fmt.Print("Masukan Umur: ")
	scanner.Scan()
	var age = scanner.Text()

	fmt.Println(name)
	fmt.Println(age)

}
