package Models

type TransactionResponse struct {
	Result []struct {
		FromAddress string `json:"from_address"`
		ToAddress   string `json:"to_address"`
		Value       string `json:"value"`
	} `json:"result"`
}
