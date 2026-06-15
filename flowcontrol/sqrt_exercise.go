package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {

	z = 1.0
	for i, diff := z, 0.0; i < 10; i++ {
		diff = (z*z - x) / (2 * z)
		z -= diff
		fmt.Printf("%v guess for Z is %v and diff is %v\n", i, z, diff)

		if math.Abs(diff) < 0.00000000000001 {
			break
		}
	}
	return z
}

func main() {
	val := 2
	fmt.Println("Lets check:\n\t", Sqrt(float64(val)), "\n ?== \n\t", math.Sqrt(float64(val)))
}
