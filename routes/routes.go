package routes

import (
	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/controllers"
	docs "github.com/alicobanserver/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routes(router *gin.Engine) {
	SetSwaggerInfo()

	router.GET("/", welcome)
	router.POST("/set", controllers.Set)
	router.GET("/get/:key", controllers.Get)
	router.POST("/flush", controllers.FlushData)
	router.POST("/getAll", controllers.GetAll)
	router.POST("/deleteSingle/:key", controllers.DeleteSingle)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Hi! I am Ali Coban.Welcome To YemekSepeti Case Study API",
	})
	return
}

// Swagger Infos
func SetSwaggerInfo() {
	docs.SwaggerInfo.Title = CONSTANTS.SWAGGERINFO_TITLE
	docs.SwaggerInfo.Description = CONSTANTS.SWAGGERINFO_DESCRIPTION
	docs.SwaggerInfo.Version = CONSTANTS.SWAGGERINFO_VERSION
	docs.SwaggerInfo.Host = CONSTANTS.SWAGGERINFO_HOST
	docs.SwaggerInfo.BasePath = CONSTANTS.SWAGGERINFO_BASEPATH
	docs.SwaggerInfo.Schemes = []string{}
}
