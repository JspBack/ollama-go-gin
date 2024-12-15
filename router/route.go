package router

import (
	"ai-thing/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine, h *handler.Handler) {
	v1 := r.Group("/api")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/api/docs/index.html")
		})

		v1.POST("/chat", h.Chat)
		v1.POST("/image", h.Image)
		v1.GET("/health", h.Health)
	}
}
