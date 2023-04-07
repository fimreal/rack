package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(r *gin.Engine) {

	swaggerBasePath := "/swagger"
	swag := r.Group(swaggerBasePath)

	swag.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
