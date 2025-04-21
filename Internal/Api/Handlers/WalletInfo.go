package Handlers

import (
	"awesomeProject/Internal/Services"
	"github.com/gin-gonic/gin"
)

type WalletRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"`
}

func GetWalletBalance(c *gin.Context) {
	var req WalletRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "address parameter is required",
		})
		return
	}

	balance, err := Services.GetWalletBalance(req.WalletAddress)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"balance": string(balance),
	})
}

func BuyEth(c *gin.Context) {

}

func SellEth(c *gin.Context) {

}
func GetWalletTransactions(c *gin.Context) {
}
