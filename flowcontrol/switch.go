package main

import (
	"fmt"
	"runtime"
	"time"
)

func osSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

}

func timeDaySwitch(t time.Weekday) {
	fmt.Println("When's ", t, "?")

	today := time.Now().Weekday()
	switch t {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In 2 days.")
	default:
		fmt.Println("Too far away :(")
	}
}

func timeSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}
}

func main() {

	osSwitch()

	timeDaySwitch(time.Tuesday)

	timeSwitch()

}
