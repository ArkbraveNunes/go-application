package router

import (
	"fmt"
	"go-basic/internal/application-openings/domain/service"

	"github.com/gin-gonic/gin"

	docs "go-basic/api"

	docsFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(router *gin.Engine) {
	service.InitServices()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	routerGroupV1 := router.Group(basePath)

	prefixOpening := "/opening"
	{
		routerGroupV1.GET(prefixOpening, service.OpeningFindHandler)
		routerGroupV1.GET(fmt.Sprintf("%s/list", prefixOpening), service.OpeningListHandler)
		routerGroupV1.POST(prefixOpening, service.OpeningCreateService)
		routerGroupV1.PUT(prefixOpening, service.OpeningUpdateService)
		routerGroupV1.DELETE(prefixOpening, service.OpeningDeleteService)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(docsFiles.Handler))
}
