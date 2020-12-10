package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nelbermora/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Inicia App")
	router.Run(":9999")
}
