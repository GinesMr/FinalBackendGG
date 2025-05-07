package Models

import "time"

type TransactionResponse struct {
	Result []struct {
		FromAddress    string    `json:"from_address"`
		ToAddress      string    `json:"to_address"`
		Value          string    `json:"value"`
		Hash           string    `json:"hash"`
		BlockTimestamp time.Time `json:"block_timestamp"`
	} `json:"result"`
}
