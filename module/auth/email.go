package auth

import (
	"bytes"
	"embed"
	"html/template"
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/components/email"
	"github.com/fimreal/rack/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:embed template/*
var emailTemplate embed.FS

var mailer email.Mailer

func (h *Handler) newToken() string {
	token := uuid.New().String()
	// h.db.UpdateUser()
	return token
}

func SetEmail(c *gin.Context) {
	if err := c.ShouldBind(&mailer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("解析邮件配置时遇到问题: %w", err)
		return
	}

	ezap.Debugf("更新邮件配置: %+v", mailer)
}

// SendVerificationEmail sends a verification email to the user.
func (h *Handler) SendVerificationEmail(c *gin.Context) {
	recipient := c.Query("recipient")

	// 检查邮箱是否存在
	exists, err := h.db.IsEmailExists(recipient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("检查邮箱是否存在时遇到问题: %w", err)
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱不存在"})
		ezap.Errorf("邮箱不存在: %s", recipient)
		return
	}

	// 读取 HTML 模板
	bodyTempl, err := emailTemplate.ReadFile("template/verification_email.html")
	if err != nil {
		ezap.Fatal("读取 HTML 模板时遇到问题: %w", err)
		return
	}
	// 解析模板
	t := template.Must(template.New("verification").Parse(string(bodyTempl)))

	verificationToken := uuid.New().String()
	verificationLink := "https://" + c.Request.Host + "/api/account/verify/email?token=" + verificationToken

	// 渲染模板
	buf := new(bytes.Buffer)
	data := map[string]string{
		"VerificationLink": verificationLink,
		"AppName":          config.AppName,
	}
	if err = t.Execute(buf, data); err != nil {
		ezap.Fatal("渲染模板时遇到问题: %w", err)
		return
	}

	// 发送邮件
	letter := &email.Letter{
		Subject:   config.AppName + " 注册账户验证",
		Recipient: recipient,
		Type:      "text/html",
		Body:      buf.String(),
	}

	err = mailer.Send(letter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("发送邮件时遇到问题: %w", err)
		return
	}
}

// SendPasswordResetEmail sends a password reset email to the user.
func (h *Handler) SendPasswordResetEmail(c *gin.Context) {
	recipient := c.Query("recipient")

	// 检查邮箱是否存在
	exists, err := h.db.IsEmailExists(recipient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("检查邮箱是否存在时遇到问题: %w", err)
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱不存在"})
		ezap.Errorf("邮箱不存在: %s", recipient)
		return
	}

	// 读取 HTML 模板
	bodyTempl, err := emailTemplate.ReadFile("template/password_reset_email.html") // 假设你有这个模板
	if err != nil {
		ezap.Errorf("读取重置密码 HTML 模板时遇到问题: %w", err)
		return
	}

	// 解析模板
	t := template.Must(template.New("passwordReset").Parse(string(bodyTempl)))

	// 生成重置密码令牌
	resetToken := uuid.New().String()
	resetLink := "https://" + c.Request.Host + "/reset/password?token=" + resetToken

	// 渲染模板
	buf := new(bytes.Buffer)
	data := map[string]string{
		"ResetLink": resetLink,
		"AppName":   config.AppName,
	}
	if err = t.Execute(buf, data); err != nil {
		ezap.Errorf("渲染重置密码模板时遇到问题: %w", err)
		return
	}

	// 发送邮件
	letter := &email.Letter{
		Subject:   config.AppName + " 重置密码请求",
		Recipient: recipient,
		Type:      "text/html",
		Body:      buf.String(),
	}

	err = mailer.Send(letter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("发送重置密码邮件时遇到问题: %w", err)
		return
	}

	// 响应成功
	c.JSON(http.StatusOK, gin.H{"message": "重置密码邮件已发送，请检查您的邮箱"})
}
