package api

import (
	"errors"
	"fmt"
)

type ConversionInformation struct {
	SourceAmount float64 `json:"sourceAmount"`
	SourceUnit   string  `json:"sourceUnit"`
	TargetAmount float64 `json:"targetAmount"`
	TargetUnit   string  `json:"targetUnit"`
	Answer       string  `json:"answer"`
}

func ConvertAmount(ingredientName string, amount float64, unit string, convertToUnit string) (*ConversionInformation, error) {
	var conversionInfo ConversionInformation
	var err error

	for i := 0; i < 3; i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/convert?apiKey=%s&ingredientName=%s&sourceAmount=%.2f&sourceUnit=%s&targetUnit=%s", API_KEY[i], ingredientName, amount, unit, convertToUnit)
		err = sendApiCall(apiUrl, &conversionInfo)

		if err != nil && err.Error() == "this api key is ratelimited" {
			continue
		} else if err != nil {
			return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
		}
		return &conversionInfo, nil
	}
	// if we did not find a good api key, throw an error (we finished looping)
	return nil, fmt.Errorf("error making request to Spoonacular API: %w", errors.New("all api keys are ratelimited"))
}
