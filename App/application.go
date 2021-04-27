package app

import (
	"github.com/gin-gonic/gin"
	_ "kasra_medrick.com/mongo/Configs/db"
	"os"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
