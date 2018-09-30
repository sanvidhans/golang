package main

import "fmt"

func main(){
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// var name = "Sanvidhan"
	//var name string

	var age int32 = 27
	size := 3.4
	//name := "Sanvidhan"
	//email := "sanvdihans@gmail.com"

	name, email := "Sanvidhan", "sanvidhans@gmail.com"


	fmt.Println(name, age, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", size)
}