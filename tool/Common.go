package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 1
	FAILED  = 0
)

func Success(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    SUCCESS,
		"message": "成功",
		"data":    data,
	})
}

func Failed(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"code":    FAILED,
		"message": "失败",
		"data":    data,
	})
}
