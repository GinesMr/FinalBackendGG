package Models

type WalletRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}
