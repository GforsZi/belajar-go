package main

import (
	"fmt"
)

var text, word string = "Hello well", "word"

func main() {
	greating := "welcome"

	var (
		a int = 50
		b int = 30
	)
	const c int = 60

	fmt.Println(a, b+c)
	fmt.Printf("%b %+d %o %#X %04d \n", a, b, a, a, b)

	fmt.Print(text, "\n")
	fmt.Print(word, " ", greating, "\n")
	fmt.Printf("my number is %v and %#v \n", a, greating)
}
