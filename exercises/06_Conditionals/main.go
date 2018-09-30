package main

import "fmt"

func main(){

	x := 6
	y := 10

	if x < y{
		fmt.Printf("%d is less than %d\n", x, y)
	} else{
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "red"

	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue"{
		fmt.Println("Color is blue")
	}else{
		fmt.Println("Color is not Red and Blue")
	}

	color = "blue"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not Red and Blue")
	}
}