package main


import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "http://www.google.com"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		os.Exit(1)
	}

	fmt.Println(response)
}
