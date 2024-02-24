package router

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
)

func InitRouter(appPort string) {
	router := gin.Default()

	initRoutes(router)

	router.Run(fmt.Sprintf(":%s", appPort))
}
