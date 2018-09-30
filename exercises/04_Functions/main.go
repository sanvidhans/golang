package main

import "fmt"


//func setName(n string){
//	name = n
//}
//func getName() string {
//	return name
//}

func greeting(name string) string{
	return "Hello "+name
}

func getSum(num1, num2 int) int{
	return num1+num2
}

func main(){
	fmt.Println(greeting("Sanvidhan"))
	fmt.Println(getSum(34,43))
}
