package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {
}

func (hc *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hc.hello)
}

// @Tags hello
// @Success 200 {string} hello
// @Router /hello [get]
func (hc *HelloController) hello(context *gin.Context) {
	context.JSON(http.StatusOK, "hello")
}
