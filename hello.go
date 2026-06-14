package main

import (
	"fmt"
)

var text, word string = "Hello well", "word"

func main() {
	greating := "welcome"

	var (
		a       uint    = 50 // uint8 uint16 uint32 uint64
		b       int     = 3  // int8 int16 int32 int64
		boolean bool    = true
		float   float32 = 3.14 // float32 float64
	)

	const c int = 60

	var array_type = [...]int{3, 6, 8, 5, 2}
	empty_array := [5]int{}
	custom_array_potition_value := [5]int{1: 2, 3: 1}

	slice_text := []string{"here", "we", "go"}
	arrslice := array_type[3:5]
	arrslice[1] = 15
	arrslice = append(arrslice, 1, 2)
	fmt.Printf("slice = %v %d \n", arrslice, len(arrslice))
	fmt.Println(len(slice_text))
	fmt.Println(cap(slice_text))

	make_slice := make([]int, 4, 5) // If the capacity parameter is not defined, it will be equal to length.
	fmt.Printf("make: %v %d %d \n", make_slice, len(make_slice), cap(make_slice))
	append_slices := append(arrslice, make_slice...)
	fmt.Printf("slices = %v \n", append_slices)

	neededArr := array_type[:len(array_type)-2]
	arrCopy := make([]int, len(neededArr))
	copy(arrCopy, neededArr)
	fmt.Printf("copy  %v \n", arrCopy)

	b |= 5
	fmt.Printf("sum is %v because b before is %03b and 5 is %03b and now b is %03b \n", b, 3, 5, b)
	b ^= 5
	fmt.Printf("sum is %v because b before is %b and 5 is %03b and now b is %03b \n", b, 7, 5, b)
	b &= 4
	fmt.Printf("sum is %v because b before is %03b and 4 is %03b and now b is %03b \n", b, 2, 4, b)

	println((b > c) || (b < c))
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
