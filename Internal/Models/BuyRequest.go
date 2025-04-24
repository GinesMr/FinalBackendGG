package Models

type BuyResponse struct {
	PrivateKey    string `json:"privateKey"`
	Amount        string `json:"amount"`
	ReciveAddress string `json:"reciveAddress"`
}
