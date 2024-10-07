package lib

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

func ProcessSum(jsonData string) (float32, float32, error) {
	var responseBody map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &responseBody)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0, 0, err
	}

	// Access account data
	account := responseBody["data"].(map[string]interface{})["account"].(map[string]interface{})

	// Calculate total value of deposits
	deposits := account["deposits"].([]interface{})
	totalDepositValue := 0.0
	for _, deposit := range deposits {
		depositMap := deposit.(map[string]interface{})
		amount, _ := strconv.ParseFloat(depositMap["amount"].(string), 64)
		decimals := depositMap["asset"].(map[string]interface{})["decimals"].(float64) // No need to convert from string
		lastPriceUSD, _ := strconv.ParseFloat(depositMap["asset"].(map[string]interface{})["lastPriceUSD"].(string), 64)
		totalDepositValue += (amount / math.Pow(10, decimals)) * lastPriceUSD
	}

	// Calculate total value of borrows
	borrows := account["borrows"].([]interface{})
	totalBorrowValue := 0.0
	for _, borrow := range borrows {
		borrowMap := borrow.(map[string]interface{})
		amount, _ := strconv.ParseFloat(borrowMap["amount"].(string), 64)
		decimals := borrowMap["asset"].(map[string]interface{})["decimals"].(float64) // No need to convert from string
		lastPriceUSD, _ := strconv.ParseFloat(borrowMap["asset"].(map[string]interface{})["lastPriceUSD"].(string), 64)
		totalBorrowValue += (amount / math.Pow(10, decimals)) * lastPriceUSD
	}

	return float32(totalDepositValue), float32(totalBorrowValue), nil
}
