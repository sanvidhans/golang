package main

import (
	"fmt"
	"math"
	"github.com/sanvidhans/exercises/03_Packages/strutil"
)

func main(){

	number := 4.5

	fmt.Println(math.Floor(number))
	fmt.Println(math.Ceil(number))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutil.Reverse("nahdivnaS"))

}
