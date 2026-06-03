package funcs

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func namedReturns(input int) (x, y int) {
	x = input + 100
	y = x / 2
	return
}

func funcs() {
	fmt.Println(add(42, 12))

	one, two := swap("Abra", "cadabra")
	fmt.Println(one, two)
	fmt.Println(namedReturns(17))
}
