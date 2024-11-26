package api

import (
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

	curr_api_key := getAPIKeys() // make sure keys are initialized

	for i := 0; i < len(curr_api_key); i++ {
		apiUrl := fmt.Sprintf("https://api.spoonacular.com/recipes/convert?apiKey=%s&ingredientName=%s&sourceAmount=%.2f&sourceUnit=%s&targetUnit=%s", curr_api_key[i], ingredientName, amount, unit, convertToUnit)
		err = sendApiCall(apiUrl, &conversionInfo)

		if err != nil {
			if err.Error() == "this api key is ratelimited" {
				continue
			} else {
				return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
			}
		}
		break
	}

	if err != nil {
		return nil, fmt.Errorf("error making request to Spoonacular API: %w", err)
	}

	return &conversionInfo, nil
}
