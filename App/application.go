package app
import "github.com/gin-gonic/gin"

var (
    router = gin.Default()
)

func StartApp(){
    mapURLs()
    router.Run(":3000")
}

