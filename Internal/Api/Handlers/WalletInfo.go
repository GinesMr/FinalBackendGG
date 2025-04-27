package Handlers

import (
	"awesomeProject/Internal/Models"
	"awesomeProject/Internal/Services"
	"github.com/gin-gonic/gin"
	"os"
	"net/url"
)

func PostWalletBalance(c *gin.Context) {
	var req Models.WalletRequest
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

func PostWalletBalancePrice(c *gin.Context) {
	var req Models.WalletRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "address parameter is required",
		})
		return
	}

	balance, err := Services.GetWalletBalancePrice(req.WalletAddress)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"balance": balance.Balance,
		"price":   balance.Price,
	})
}

func SendEth(c *gin.Context) {
	var res Models.BuyResponse
	err := c.ShouldBindJSON(&res)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "One or more parameters are missing",
		})
	}
	err = Services.SendEthFun(res.ReciveAddress, res.PrivateKey, res.Amount)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	if err == nil {
		c.JSON(200, gin.H{
			"transaction": "success",
		})
	}

}

func SellEth(c *gin.Context){
c.JSON(200,gin.H{
	"PedroSanchez": "Hijo de puta",
	"Liebe": "Es maricon",
})
}

func GetWalletTransactions(c *gin.Context) {
}

func EthPrice(c *gin.Context) {
	res, err := Services.EthPriceService()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"price": res.Result[0].UsdPrice,
	})
}
//I should to put this piece of code in another part like Services
func UrlRampCreator(c*gin.Context){
WalletAddressRamp := url.URL{
        Scheme: "https",
        Host: os.Getenv("RAMP_BASE_URl"),
			}
			c.JSON(200,gin.H{
        "url":WalletAddressRamp.String(),
			})
	}
