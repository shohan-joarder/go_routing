package main

import (
	"fmt"
	"log"
	"net/http"
)

func fromHandler(w http.ResponseWriter, r *http.Request)  {
	if err :=r.ParseForm(); err !=nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	
	fmt.Fprintf(w, "Post request successfull");
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	if(r.URL.Path !="hello"){
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if(r.Method != "GET"){
		http.Error(w,"Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main()  {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/from", fromHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Start server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}

}