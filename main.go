package main

import (
	"fmt"
	"net/http"
	"strconv"
	"url-shortener/handler"
	"url-shortener/helper"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Hello Go URL Shortener",
		})
	})

	router.POST("/", handler.CreateShortUrl)
	router.GET("/:short_url", handler.HandleShortUrlRedirect)

	store.InitializeStore()

	port, _ := strconv.Atoi(helper.GetEnvVariableKey("APP_PORT", "8080"))
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server: %v", err))
	}
}
