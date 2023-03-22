package handler

import (
	"net/http"
	"url-shortener/helper"
	"url-shortener/shortener"
	"url-shortener/store"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(ctx *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.InitialUrl, creationRequest.UserID)
	store.SaveUrlMapping(shortUrl, creationRequest.InitialUrl, creationRequest.UserID)

	host := helper.GetEnvVariableKey("APP_URL", "http://localhost:8080/")

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("short_url")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}
