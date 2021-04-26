package app

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Configs/db"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	port := db.DotEnv("PORT")
	router.Run(":" + port)
}
