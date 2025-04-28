package Services

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
)

func UrlRampCreator(c *gin.Context) {
	REAL_URL := os.Getenv("RAMP_BASE_URl") + "APIKEY"
	WalletAddressRamp := url.URL{
		Scheme: "https",
		Host:   REAL_URL,
	}
	c.JSON(200, gin.H{
		"url": WalletAddressRamp.String(),
	})
}
