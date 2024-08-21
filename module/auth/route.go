package auth

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module/auth/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("all_services") {
		return
	}

	jwt.SetJWTSecret(viper.GetString("auth_seed"))
	jwt.SetJWTSigningMethod(viper.GetString("auth_method"))

	// 创建路由组 /api
	r := g.Group(RoutePrefix)

	// skipPaths 不需要权限验证接口
	skipPaths := []string{"/api/account/login", "/api/account/signin", "/api/account/token"}
	r.Use(jwt.JWTMiddleware(skipPaths))

	// 路由组 /api/account
	addAccountRoutes(r)
}

func addAccountRoutes(g *gin.RouterGroup) {
	// docs.SwaggerInfo.BasePath = "/account"
	// account := g.Group(docs.SwaggerInfo.BasePath)
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	account := g.Group("/account")
	h, err := NewHandler()
	if err != nil {
		ezap.Fatal(err)
	}
	account.POST("/signin", h.Signin)
	account.POST("/login", h.Login)

	account.POST("/token/validate", h.IsValidToken)
	account.POST("/token/renew", h.RenewToken)
}
