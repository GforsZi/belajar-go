package main

import (
	"fmt"
)

var text, word string = "Hello well", "word"

func main() {
	greating := "welcome"

	var (
		a       uint    = 50 // uint8 uint16 uint32 uint64
		b       int     = 30 // int8 int16 int32 int64
		boolean bool    = true
		float   float32 = 3.14 // float32 float64
	)

	const c int = 60

	var array_type = [...]int{3, 6, 8, 5, 2}
	empty_array := [5]int{}
	custom_array_potition_value := [5]int{1: 2, 3: 1}

	if boolean {
		array_type[0] = 1
		fmt.Println(len(array_type))
		fmt.Println(array_type)
		fmt.Println(empty_array)
		fmt.Println(custom_array_potition_value)
	}

	fmt.Println(a, b+c, float*float)
	fmt.Printf("%b %+d %o %#X %04d \n", a, b, a, a, b)

	fmt.Print(text, "\n")
	fmt.Print(word, " ", greating, "\n")
	fmt.Printf("my number is %v and %#v type %T %%\n", a, greating, greating)
}
