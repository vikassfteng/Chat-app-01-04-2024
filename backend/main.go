package main

import "net/http"

func serverWs(w http.ResponseWriter, r *http.Request) {
	// Print when a new client connects

	println("Client connected")
}
func main() {
	// Call the function that will start the server
	http.ListenAndServe(":9000", nil)

}
