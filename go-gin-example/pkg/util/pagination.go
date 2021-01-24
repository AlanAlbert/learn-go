package util

import (
	"github.com/anhoder/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(context *gin.Context) int {
	result := 0
	page, _ := com.StrTo(context.Query("page")).Int()

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
