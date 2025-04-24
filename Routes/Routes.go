package Routes

import (
	"awesomeProject/Endpoints"
	"awesomeProject/Internal/Api/Handlers"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	routes := gin.Default()

	routesComplete := routes.Group("/api/v1")
	{
		routesComplete.POST(Endpoints.GetWalletBalanceEndpoint, Handlers.PostWalletBalance)
		routesComplete.POST(Endpoints.SendEth, Handlers.SendEth)
		routesComplete.POST(Endpoints.SellEth, Handlers.SellEth)
		routesComplete.POST(Endpoints.GetWalletBalancePriceEndpoint, Handlers.PostWalletBalancePrice)
	}

	return routes
}
