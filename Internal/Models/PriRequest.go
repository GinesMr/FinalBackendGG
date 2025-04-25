package Models

type PriRequest struct {
	Result []struct {
		USDPrice float64 `json:"usd_price"`
	} `json:"result"`
}
