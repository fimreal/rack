package email

import (
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
)

type MailerInterface interface {
	SendMail(c *gin.Context)
}

func SendMail(c *gin.Context) {
	var letter Letter
	if err := c.ShouldBind(&letter); err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ezap.Debugf("请求发送邮件, 传入收件人: %v, 发送标题: %s, 内容: %s", letter.Mailto, letter.Subject, letter.Body)

	if err := Mailto(&letter); err != nil {
		ezap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ezap.Debug("邮件发送成功")
	c.JSON(http.StatusOK, gin.H{"result": "邮件发送成功"})
}
