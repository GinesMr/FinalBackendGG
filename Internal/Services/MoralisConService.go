package Services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type WalletBalance struct {
	Balance string `json:"balance"`
}

func GetWalletBalance(ad string) (string, error) {
	API_KEY := os.Getenv("API_KEY")
	BASE_URL := os.Getenv("URL")
	URL2 := os.Getenv("URL2")

	if !strings.HasPrefix(BASE_URL, "https://") {
		BASE_URL = "https://" + BASE_URL
	}

	REAL_URL := BASE_URL + ad + URL2

	if _, err := url.Parse(REAL_URL); err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	req, err := http.NewRequest("GET", REAL_URL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	var balance WalletBalance
	if err := json.NewDecoder(res.Body).Decode(&balance); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return balance.Balance, nil
}
