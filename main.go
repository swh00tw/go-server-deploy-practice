package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Serving on", port)
	http.ListenAndServe(":"+port, nil)
}
