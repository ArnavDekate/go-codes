package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL of the API endpoint
	apiUrl := "https://api.example.com/data"

	//send SET request
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error making GET request : %s\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body : %s\n", err)
		return
	}

	println(string(body))
}
