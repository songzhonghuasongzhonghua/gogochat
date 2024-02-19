package tool

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs2 "gogochat/docs"
)

func InitSwagger(engine *gin.Engine) {
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8090/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	docs2.SwaggerInfo.BasePath = ""

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
