package Services

import (
	"awesomeProject/Internal/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func NftCollection(WalletAddres string) (Models.NftResponseModel, error) {
	API_KEY := os.Getenv("API_KEY")
	URL := os.Getenv("URL") + WalletAddres
	URL_2 := os.Getenv("NFT_URL2")
	REAL_URL := URL + URL_2

	req, _ := http.NewRequest("GET", REAL_URL, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", API_KEY)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return Models.NftResponseModel{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	var NftCollection Models.NftResponseModel
	if err := json.NewDecoder(res.Body).Decode(&NftCollection); err != nil {
		return Models.NftResponseModel{}, fmt.Errorf("error decoding response: %w", err)
	}
	return NftCollection, nil

}
