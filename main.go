package main


import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	url := "http://www.google.com"

	httpResponse, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP error:", err)
		os.Exit(1)
	}

	// bytes := make([]byte, 99999)
	// httpResponse.Body.Read(bytes)
	// response := string(bytes)
	// fmt.Println("response: ",response)

	// io.Copy(os.Stdout, httpResponse.Body)

	htmlWriter := NewHTMLWriter("response.html")
	_, err = io.Copy(htmlWriter, httpResponse.Body)
	if err != nil {
		fmt.Println("io.Copy error: ", err)
	}
}
