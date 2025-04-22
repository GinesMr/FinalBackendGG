package Models

type WalletRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

type PriRequest struct {
	Precio int `json:"precio" binding:"required"`
}
