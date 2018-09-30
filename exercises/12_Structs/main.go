package main

import (
	"strconv"
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	city string
	phone string
	gender string
	age int
	email string
}

func (p Person) greet() string{
	return "Hello I am "+p.firstName+" "+p.lastName+" and I am "+strconv.Itoa(p.age)+" years old."
}

func (p *Person) hasBirthday(){
	p.age++
}

func (p *Person) getMarried(spouseLastName string){
	if p.gender == "m"{
		return
	}else{
		p.lastName = spouseLastName
	}
}

func main(){
	//Init person using struct

	//person1 := Person{firstName:"Sanvidhan", lastName:"Sonone", city:"Pune", phone:"8007063030", email:"sanvidhans@gmail.com", age:27}

	//fmt.Println(person1)
	//alternative
	person1 := Person{"Brad", "Mokan", "New York", "7845875485", "m",27, "brad@gmail.com"}
	person2 := Person{"Sreylly", "Sohan", "New York", "9856485652", "f",23, "shrelly@gmail.com"}
	//fmt.Println(person1)
	//fmt.Println(person1.firstName)
	//person1.age++
	//fmt.Println(person1.age)
	person1.hasBirthday()
	person1.getMarried("Sohan")

	person2.getMarried("Mokan")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
