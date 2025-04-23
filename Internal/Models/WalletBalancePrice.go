package Models

type WalletBalancePrice struct {
	Balance string `json:"balance_formatted"`
	Price   string `json:"usd_price"`
}
