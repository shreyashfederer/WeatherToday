package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	httpclient := http.Client{}

	nameOfCity := "mumbai"

	reqURL := fmt.Sprintf("https://www.metaweather.com/api/location/search?query=%s", nameOfCity)

	fmt.Println(reqURL)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)

	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)

	}

	req.Header.Set("accept", "application/json")

	response, err := httpclient.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)

	}
	defer response.Body.Close()

	fmt.Println(response.Body)
}
