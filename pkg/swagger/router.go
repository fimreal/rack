package swagger

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(r *gin.RouterGroup) {
	if !viper.GetBool("swagger") {
		return
	}

	swaggerBasePath := "/swagger"
	swag := r.Group(swaggerBasePath)

	swag.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
