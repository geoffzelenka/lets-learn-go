package vars

import (
	"fmt"
	"math"
	"math/cmplx"
)

var d, pascal, haskell bool

var i, j int = 1, 2

const LilPi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func shortDecls() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "No!"

	fmt.Println(i, j, k, c, python, java)
}

func basicTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}

func conversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = f // compiler error
	var z uint = uint(f)

	fmt.Println(x, y, z)

}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func vars() {
	var i int

	var c, python, java = true, false, "no!"
	fmt.Println(i, d, pascal, haskell)
	fmt.Println(i, j, c, python, java)

	shortDecls()

	basicTypes()

	conversions()

	fmt.Println("Time to eat ", LilPi)
	//LilPi = 3.1416 //obvi compiler error
	fmt.Printf("LilPi is a %T with value %v\n", LilPi, LilPi)

	numericConstants()
}
