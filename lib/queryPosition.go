package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func QueryPosition() (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	apiToken := os.Getenv("THE_GRAPH_API_KEY")
	if apiToken == "" {
		panic("No Graph api token available")
	}
	ethAccount := os.Getenv("ETH_ACCOUNT")
	if ethAccount == "" {
		panic("No Graph api token available")
	}

	url := fmt.Sprintf(`https://gateway.thegraph.com/api/%s/subgraphs/id/4xyasjQeREe7PxnF6wVdobZvCw5mhoHZq3T7guRpuNPf`, apiToken)

	// Construct the GraphQL query
	query := fmt.Sprintf(`{
		account(id: "%s") {
			deposits {
				id,
				amount,
				asset {
					id,
					lastPriceUSD,
					name,
					decimals
				}
			},
			borrows {
				id,
				amount,
				asset {
					id,
					lastPriceUSD,
					name,
					decimals
				}
			}
		}
	}`, ethAccount)

	payload := map[string]interface{}{
		"query": query,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return "", err
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		fmt.Println("Error decoding response:", err)
		return "", err
	}

	// Print the response as formatted JSON
	formattedJSON, err := json.MarshalIndent(responseBody, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return "", err
	}

	return string(formattedJSON), nil
}
