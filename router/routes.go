package router

import (
	"application-openings/handler"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	handler.InitHandler()
	routerGroupV1 := router.Group("/api/v1")

	routePrefixOpening := "/opening"
	{
		routerGroupV1.GET(routePrefixOpening, handler.FindOpeningHandler)
		routerGroupV1.GET(routePrefixOpening+"/list", handler.ListOpeningHandler)
		routerGroupV1.POST(routePrefixOpening, handler.CreateOpeningHandler)
		routerGroupV1.PUT(routePrefixOpening, handler.UpdateOpeningHandler)
		routerGroupV1.DELETE(routePrefixOpening, handler.DeleteOpeningHandler)
	}

}
