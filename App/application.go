package app

import (
	"github.com/gin-gonic/gin"
	_ "kasra_medrick.com/mongo/Configs/db"
 )

var (
	router = gin.Default()
)

func StartApp() {
	mapURLs()
	// port := db.DotEnv("PORT")
	// router.Run(":" + port)
	router.Run(":8080")

}
