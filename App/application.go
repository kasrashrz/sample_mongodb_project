package app

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Configs/db"
	_ "kasra_medrick.com/mongo/Configs/db"

	)

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	port := db.ViperConfigVariable("PORT")
	router.Run(":" + port)
}
