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
		routesComplete.GET(Endpoints.GetWalletBalanceEndpoint, Handlers.GetWalletBalance)
		routesComplete.POST(Endpoints.PostWalletAddres, Handlers.PostWalletAddres)
	}

	return routes
}
