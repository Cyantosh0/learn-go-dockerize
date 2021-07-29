package route

import (
	"github.com/gin-gonic/gin"

	"github.com/cyantosh0/dockerize-go-app/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/apple/b", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Up and running"})
		})
		user.GET("/", controller.GetUsers)
		user.POST("/", controller.CreateUser)
		user.GET("/:id", controller.GetUserByID)
	}
	return r
}
