package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ConversionInformation struct {
	SourceAmount float64 `json:"sourceAmount"`
	SourceUnit   string  `json:"sourceUnit"`
	TargetAmount float64 `json:"targetAmount"`
	TargetUnit   string  `json:"targetUnit"`
	Answer       string  `json:"answer"`
}

func ConvertAmount(ingredientName string, amount float64, unit string, convertToUnit string) (*ConversionInformation, error) {
	var resp *http.Response
	var err error

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/convert?apiKey=%s&ingredientName=%s&sourceAmount=%.2f&sourceUnit=%s&targetUnit=%s", API_KEY[i], ingredientName, amount, unit, convertToUnit)
		resp, err = http.Get(apiUrl)

		if err != nil && err.Error() == "this api key is ratelimited" {
			continue
		} else {
			if err != nil {
				return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				if resp.StatusCode == 402 || resp.StatusCode == 429 {
					return nil, fmt.Errorf("this api key is ratelimited")
				}
				return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			var conversionInfo ConversionInformation
			err = json.Unmarshal(body, &conversionInfo)
			if err != nil {
				return nil, fmt.Errorf("error parsing JSON: %w", err)
			}

			return &conversionInfo, nil
		}
	}
	// if we did not find a good api key, throw an error (we finished looping)
	return nil, fmt.Errorf("error making request to Spoonacular API: %w", errors.New("all api keys are ratelimited"))
}
