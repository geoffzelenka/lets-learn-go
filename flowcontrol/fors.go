package main

import (
	"fmt"
	"math"
	"math/rand"
)

func yolo() bool {
	return !(rand.Int31n(100) == 42)

}

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i * i
	}

	fmt.Println(sum)
	//fmt.Println(i) // not defined, went out of scope after the for loop

	sum = 1

	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}

	fmt.Println(sum)

	count := 0

	for yolo() {
		count++
	}

	fmt.Printf("Took %v tries to get out of the yolo\n", count)

	for {
		//ever
		if !yolo() {
			break
		}
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
