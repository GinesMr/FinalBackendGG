package Models

type NftResponseModel struct {
	Result []struct {
		TokenAddress string `json:"token_address"`
		ContractType string `json:"contract_type"`
		Name         string `json:"name"`
	} `json:"result"`
}
