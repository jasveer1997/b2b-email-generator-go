package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := 8080

	http.HandleFunc("/hello", HelloHandler)
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello, World!")
}
