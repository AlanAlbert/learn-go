package routers

import (
	"github.com/anhoder/go-gin-example/pkg/setting"
	v1 "github.com/anhoder/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(setting.AppMode)

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiV1 := engine.Group("/api/v1")

	// Tag标签
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.UpdateTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return engine
}