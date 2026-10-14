package main

import(
	"fmt"
	"net/http"
)

// When router mux.HandleFunc is called -> It will create a route "/hello" and call helloHandler(). 
// Inside helloHandler() we have two arguments w -> Writer, r -> Request
// w will write the response ("Hello World") to the client (UI).
// r -> Request is used to read the request from the client (Not used here). e.g. login credentials username, password input from user from UI.
// For handling the request we pass this mux to the ListenAndServe function. 
func helloHandler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintln(w, "Hello World") // this is the response.
}

func aboutHandler(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprintln(w, "I am Shajib. I am a Software Engineer. Now I am learning Go")
}


func main() {
	mux := http.NewServeMux() // mux is a Router.

	mux.HandleFunc("/hello", helloHandler) // here pattern "/hello" is a Route.

	mux.HandleFunc("/about", aboutHandler) // route

	fmt.Println("Server running on port: 3000") // starting a local server at port 3000. 

	err := http.ListenAndServe(":3000", mux)//Listen for incoming requests if server port is free.

	if err != nil {
		fmt.Println("Error starting the server", err) 
	}
}