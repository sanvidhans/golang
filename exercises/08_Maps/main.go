package main

import "fmt"

func main(){
	// Define maps
	//emails := make(map[string]string)

	//emails["abc"] = "abc@gmail.com"
	//emails["bca"] = "bca@gmail.com"
	//emails["pqr"] = "pqr@gmail.com"

	emails := map[string]string{"abc":"abc@gmail.com","bca":"bca@gmail.com","pqr":"pqr@gmail.com",}

	fmt.Println(emails)
	fmt.Println(emails["abc"])

	//delete bca email from map
	delete(emails, "abc")
	fmt.Println(emails)
}
