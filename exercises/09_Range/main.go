package main

import "fmt"

func main(){
	ids := []int{33, 76, 54, 23, 11, 2}

	//Print with index
	for i, id := range ids{
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//print without index
	for _, id := range ids{
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids{
		sum += id
	}

	fmt.Println("Sum", sum)

	//Range with map
	emails := map[string]string{"Bob":"bob@gmail.com", "Jak":"jak@gmail.com", "Sharon":"sharon@gmail.com"}

	for k, v := range emails{
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails{
		fmt.Println("Name:"+k)
	}
}
