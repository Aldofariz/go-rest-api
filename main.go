package main

import (
	"fmt"
	"gin-socmed/config"
	"gin-socmed/docs"
	"gin-socmed/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api
// @title GIN-SOCMED API
// @version 1.0
// @description Ini adalah dokumentasi REST API untuk project Gin-Socmed
// @host localhost:8080
// @termsOfService  http://swagger.io/terms/

func main() {

	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// HelloHandler godoc
	// @Summary Menyapa pengguna
	// @Description Mengembalikan pesan hello
	// @Tags hello
	// @Accept json
	// @Produce json
	// @Success 200 {string} string "Hello, world!"
	// @Router /hello [get]
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)
	router.PostRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
