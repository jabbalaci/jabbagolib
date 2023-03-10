package main

import (
	"fmt"

	"github.com/jabbalaci/jabbagolib/jweb"
)

func main() {
	url := "http://example.com"

	html, err := jweb.FetchHTML(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(html)
}
