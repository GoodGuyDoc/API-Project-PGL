package api_retrieval

import (
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
