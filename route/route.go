package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/cyantosh0/dockerize-go-app/controller"
	"github.com/cyantosh0/dockerize-go-app/middleware"
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
		user.POST("/", controller.CreateUser)
		user.Use(middleware.AuthRequired)
		{
			user.GET("/", controller.GetUsers)
			user.GET("/:id", controller.GetUserByID)
		}
	}
	return r
}
