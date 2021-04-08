package app

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Configs"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	port := Configs.DotEnv("PORT")
	router.Run(":" + port)
}
