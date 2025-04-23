package Models

type WalletResponse struct {
	Result []struct {
		BalanceFormatted string  `json:"balance_formatted"`
		UsdPrice         float64 `json:"usd_price"`
	} `json:"result"`
}
