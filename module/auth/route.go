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
	r := g.Group("/account")
	h, err := NewHandler()
	if err != nil {
		ezap.Fatal(err)
	}
	r.POST("/signin", h.Signin)
	r.POST("/login", h.Login)

	r.POST("/token/validate", h.IsValidToken)
	r.POST("/token/renew", h.RenewToken)
}
