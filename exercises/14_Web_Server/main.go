package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"<h1>Hey, Welcome to Web Server!</h1>")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>It's a About Page!</h1>")
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Stating...")
	fmt.Println("Server Started!")
	fmt.Println("Open brower and enter URL: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}