package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ConversionInformation struct {
	SourceAmount float64 `json:"sourceAmount"`
	SourceUnit   string  `json:"title"`
	TargetAmount float64 `json:"image"`
	TargetUnit   string  `json:"analyzedInstructions"`
	Answer       string  `json:"extendedIngredients"`
}

func ConvertAmount(ingredientName string, amount float64, unit string, convertToUnit string) (*ConversionInformation, error) {
	var resp *http.Response
	var err error

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/convert?apiKey=%s&ingredientName=%s&sourceAmount=%.2f&sourceUnit=%s&targetUnit=%s", API_KEY[i], ingredientName, amount, unit, convertToUnit)
		resp, err = http.Get(apiUrl)

		if err.Error() == "this api key is ratelimited" {
			continue
		} else {
			break
		}

	}
	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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
