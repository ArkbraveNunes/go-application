package router

import (
	"application-openings/handler"

	"github.com/gin-gonic/gin"

	docs "application-openings/docs"

	docsFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(router *gin.Engine) {
	handler.InitHandler()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	routerGroupV1 := router.Group(basePath)

	routePrefixOpening := "/opening"
	{
		routerGroupV1.GET(routePrefixOpening, handler.FindOpeningHandler)
		routerGroupV1.GET(routePrefixOpening+"/list", handler.ListOpeningHandler)
		routerGroupV1.POST(routePrefixOpening, handler.CreateOpeningHandler)
		routerGroupV1.PUT(routePrefixOpening, handler.UpdateOpeningHandler)
		routerGroupV1.DELETE(routePrefixOpening, handler.DeleteOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(docsFiles.Handler))

}
