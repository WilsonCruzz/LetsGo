package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var x int = 3
	fmt.Println(x)
	fmt.Scanln(&x)
	fmt.Println(x * 2)

}

// coding > build > execute
// go build name.go
