package routes

import (
	"net/http"
	"ruizhiyuan/controller"
	"ruizhiyuan/logger"
	"ruizhiyuan/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "http server ok!",
		})
	})

	r.POST("/login", controller.LoginHandler)

	admin := r.Group("/admin", middleware.JWTAuthMiddleware())
	{
		admin.POST("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "login success",
			})
		})
	}

	return r
}
