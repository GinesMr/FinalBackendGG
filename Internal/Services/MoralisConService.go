package Services

import (
	"awesomeProject/Internal/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

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

	var balance Models.WalletBalance
	if err := json.NewDecoder(res.Body).Decode(&balance); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return balance.Balance, nil
}

func GetWalletBalancePrice(ad string) (Models.WalletBalancePrice, error) {
	API_KEY := os.Getenv("API_KEY")
	BASE_URL := os.Getenv("URL4")
	URL3 := os.Getenv("URL3")

	if !strings.HasPrefix(BASE_URL, "https://") {
		BASE_URL = "https://" + BASE_URL
	}

	REAL_URL := BASE_URL + ad + URL3

	if _, err := url.Parse(REAL_URL); err != nil {
		return Models.WalletBalancePrice{}, fmt.Errorf("invalid URL: %w", err)
	}

	req, err := http.NewRequest("GET", REAL_URL, nil)
	if err != nil {
		return Models.WalletBalancePrice{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Models.WalletBalancePrice{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	var response Models.WalletResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Models.WalletBalancePrice{}, fmt.Errorf("error decoding response: %w", err)
	}
	balanceFloat, parseErr := strconv.ParseFloat(response.Result[0].BalanceFormatted, 64)
	if parseErr != nil {
		return Models.WalletBalancePrice{}, fmt.Errorf("error parsing balance: %w", parseErr)
	}
	price := response.Result[0].UsdPrice * balanceFloat
	balance := Models.WalletBalancePrice{
		Balance: response.Result[0].BalanceFormatted,
		Price:   fmt.Sprintf("%.2f", price),
	}
	return balance, nil
}
