package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/cyantosh0/dockerize-go-app/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	}))

	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Up and running....."})
	})

	user := r.Group("/user")
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "User routes Up and running....."})
		})

		user.GET("/get-users", controller.GetUsers)
		user.POST("/create", controller.CreateUser)
		user.GET("/get-user/:id", controller.GetUserByID)
	}
	return r
}
