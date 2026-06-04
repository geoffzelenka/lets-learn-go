package hello

import (
	"fmt"
	"math/rand"
)

func Hello() {
	fmt.Println("Hello world! Lets get random: ", rand.Intn(10), " have fun!")
}
