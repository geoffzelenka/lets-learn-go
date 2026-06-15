package main

import "fmt"

func main() {
	defer fmt.Println("world!")

	fmt.Println("Hello")

	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
