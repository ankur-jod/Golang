package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter a number :")
	var a int
	fmt.Scanln(&a)
	fmt.Println(a)

	var n string
	fmt.Scan(&n) //It will not take "Ankur Mukherjee " it breaks in the spaces and print ankur
	fmt.Println(n)

	var an string
	fmt.Scan(&an)
	fmt.Println(an)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println(name)

}
