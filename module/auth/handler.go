package auth

import (
	"net/http"
	"time"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/module/auth/jwt"
	"github.com/fimreal/rack/module/auth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	db models.DBLayer
}

type HandlerInterface interface {
	Signin(c *gin.Context)
	Login(c *gin.Context)
	IsValidToken(c *gin.Context)
	RenewToken(c *gin.Context)
}

func NewHandler() (HandlerInterface, error) {
	db, err := models.NewORM()
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrateUser()
	return &Handler{
		db: db,
	}, err
}

func (h *Handler) Signin(c *gin.Context) {
	var userData models.AddUserData
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("解析用户数据时遇到问题: %w", err)
		return
	}

	// 检查用户是否已经存在，并获取冲突字段
	conflictField, err := h.db.CheckUserExistence(userData.Username, userData.Email, userData.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db err"})
		return
	}
	if conflictField != "" {
		c.JSON(http.StatusConflict, gin.H{
			"error": conflictField + " 已被注册",
		})
		return
	}

	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "加密密码时遇到问题"})
		return
	}
	user := models.User{
		Username:    userData.Username,
		Password:    string(encodedPassword),
		Email:       models.Str2NullStr(userData.Email),
		PhoneNumber: models.Str2NullStr(userData.PhoneNumber),
		Nickname:    userData.Username,
	}

	if err := h.db.CreateUser(&user); err != nil { // 假设 AddUser 方法已实现
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户时遇到问题"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Login 用户登录处理
func (h *Handler) Login(c *gin.Context) {
	var loginData models.LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ezap.Errorf("解析用户数据时遇到问题: %w", err)
		return
	}

	// 获取用户
	user, err := h.db.GetUser(loginData.UserIdentifier)
	if err != nil {
		ezap.Errorf("获取用户[%s]时遇到问题: %w", loginData.UserIdentifier, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "账号或密码错误"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "账号或密码错误"})
		return
	}

	// 生成 JWT
	accessToken, refreshToken, err := jwt.CreateTokens(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 遇到错误"})
		ezap.Errorf("生成 token 时遇到问题: %w", err)
		return
	}

	loginResult := models.LoginResult{
		ID:          user.ID,
		Username:    user.Username,
		Email:       models.NullStr2Str(user.Email),
		PhoneNumber: models.NullStr2Str(user.PhoneNumber),
		Nickname:    user.Nickname,
		Role:        user.Role,

		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// 返回 token 和用户信息
	c.JSON(http.StatusOK, loginResult)
}

// IsValidToken 验证 token 是否有效
func (h *Handler) IsValidToken(c *gin.Context) {
	// 从请求头中提取 token
	tokenString := c.Request.Header.Get("Authorization")

	// 确保 Authorization 字段存在
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Bearer Token 格式: "Bearer <token>"
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:] // 去掉 "Bearer "
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}

	valid, claims, err := jwt.DecodeToken(tokenString, "access"+jwt.Secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 accessToken"})
		ezap.Errorf("解码 accessToken 时遇到问题: %s", err)
		return
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 accessToken"})
		ezap.Errorf("accessToken 无效")
		return
	}

	// 可以从 claims 中获取用户相关信息
	id := (*claims)["id"] // 假设你在生成 token 时包含了 user_id, 注意 json 的数字在 golang 中默认时 float64 类型

	// 返回用户信息或进行其他业务逻辑
	c.JSON(http.StatusOK, gin.H{"message": "accessToken 有效", "id": id})
}

// RenewToken 刷新 token
func (h *Handler) RenewToken(c *gin.Context) {
	var refreshTokenData struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	// 绑定 JSON 数据
	if err := c.ShouldBindJSON(&refreshTokenData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的输入"})
		ezap.Errorf("解析用户数据时遇到问题: %s", err)
		return
	}

	// 解码 refreshToken
	refreshSecret := "refresh" + jwt.Secret
	valid, claims, err := jwt.DecodeToken(refreshTokenData.RefreshToken, refreshSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 refresh_token"})
		ezap.Errorf("解码 token 时遇到问题: %s", err)
		return
	}

	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "无效的 refresh_token"})
		ezap.Errorf("token 无效")
		return
	}

	id := (*claims)["id"].(float64)
	accessSecret := "access" + jwt.Secret
	exp := time.Now().Add(jwt.ExpTime).Unix()
	accessToken, _ := jwt.CreateToken(uint(id), exp, accessSecret)
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
