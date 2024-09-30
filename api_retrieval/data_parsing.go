package api_retrieval

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Sends an HTTP Get request to the spoonacular api with the specified request string. Parses the response and returns the byte array.
func getSpoonData(request string) (data []byte, err error) {
	resp, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 402 {
		return nil, errors.New("out of API requests for the day")
	} else if resp.StatusCode == 429 {
		return nil, errors.New("slow down! too many requests. Try again later")
	}

	pointsUsed := resp.Header.Get("X-API-Quota-Request")
	pointsLeft := resp.Header.Get("X-API-Quota-Left")
	fmt.Printf("This search used %s points. There are %s points left today.\n", pointsUsed, pointsLeft)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
