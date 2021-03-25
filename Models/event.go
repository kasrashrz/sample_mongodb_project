package Models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var client *mongo.Client
type  EventEndTimeType uint
const (
	absolute EventEndTimeType = iota
	relative
)
type Event struct {
	Name string
	Market_name string
	Env string
    EventEndTime EventEndTimeType
	//states map[]
	Repetition Repetition
}

type Event_Handler interface {
	CREATE(ctx *gin.Context)
	RETRIEVE(ctx *gin.Context)
}
var userModel = new(Event)

func (u Event) RETRIEVE(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		user, err := userModel.GetByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": user})
		return
	}
	fmt.Println(c.Param("id"))
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}