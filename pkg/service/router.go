package service

import (
	"github.com/fimreal/rack/pkg/service/aliyun"
	"github.com/fimreal/rack/pkg/service/email"
	"github.com/fimreal/rack/pkg/service/randomstring"
	"github.com/gin-gonic/gin"
)

func AddServiceRoutes(r *gin.RouterGroup) {

	serviceBasePath := "/s"
	srv := r.Group(serviceBasePath)

	// 阿里云
	srv.POST("/addsgrule", aliyun.Allow) // 添加安全组准入规则
	// 电子邮件
	srv.POST("/mailto", email.SendMail) //  邮件发送服务
	// other
	srv.POST("/random", randomstring.GenRandomString)
}
