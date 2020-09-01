package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./assets")))
	if err != nil {
		fmt.Println("Failed to start calculator server", err)
		os.Exit(1)
	}
}
