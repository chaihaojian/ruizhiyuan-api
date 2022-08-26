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

	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors())

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
		article := admin.Group("/article")
		{
			article.GET("/all", controller.GetAllArticleHandler)
			article.POST("/add", controller.AddArticleHandler)
			article.POST("/delete")
			article.POST("/update")

		}
		video := admin.Group("/video")
		{
			video.GET("/all", controller.GetAllVideoHandler)
			video.POST("/add", controller.AddVideoHandler)
			video.DELETE("/delete", controller.DeleteVideoHandler)
			video.POST("/update")
		}
		material := admin.Group("/material")
		{
			material.GET("/all", controller.GetAllMaterialHandler)
			material.POST("/add", controller.AddMaterialHandler)
			//material.DELETE("/delete", controller.DeleteMaterialHandler)
			//material.POST("/update")
		}
	}

	return r
}
