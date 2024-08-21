package jwt

import (
	"fmt"
	"time"

	"github.com/fimreal/goutils/ezap"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	ExpTime       = 2 * time.Minute
	RefreshTime   = 14 * 24 * time.Hour
	Secret        string
	SigningMethod jwt.SigningMethod
)

func SetJWTSecret(secret string) {
	Secret = secret
	if Secret == "" {
		Secret = uuid.New().String()
		ezap.Warn("未配置 auth secret 值, 自动生成(仅用作测试环境)使用: ", Secret)
	}
}

func SetJWTSigningMethod(method string) {
	switch method {
	// HMAC 最快，较为安全
	case "HS256":
		SigningMethod = jwt.SigningMethodHS256
	case "HS384":
		SigningMethod = jwt.SigningMethodHS384
	case "HS512":
		SigningMethod = jwt.SigningMethodHS512
	// RSA 安全性差，速度慢，不推荐使用
	case "RS256":
		SigningMethod = jwt.SigningMethodRS256
	case "RS384":
		SigningMethod = jwt.SigningMethodRS384
	case "RS512":
		SigningMethod = jwt.SigningMethodRS512
	// ECDSA 短且安全，但是计算量大
	case "ES256":
		SigningMethod = jwt.SigningMethodES256
	case "ES384":
		SigningMethod = jwt.SigningMethodES384
	case "ES512":
		SigningMethod = jwt.SigningMethodES512
	case "PS256":
		SigningMethod = jwt.SigningMethodPS256
	// 综合选择
	default:
		SigningMethod = jwt.SigningMethodHS256
	}
	ezap.Debugf("jwt 签名算法: %s", SigningMethod.Alg())
}

func CreateToken(id uint, expire int64, secret string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["expire"] = expire
	at := jwt.NewWithClaims(SigningMethod, atClaims)
	return at.SignedString([]byte(secret))
}

func CreateTokens(id uint) (string, string, error) {

	expire := time.Now().Add(ExpTime).Unix()
	accessToken, err := CreateToken(id, expire, "access"+Secret)
	if err != nil {
		return "", "", err
	}
	refresh := time.Now().Add(RefreshTime).Unix()
	refreshToken, err := CreateToken(id, refresh, "refresh"+Secret)

	return accessToken, refreshToken, err
}

func DecodeToken(tokenString, secret string) (bool, *jwt.MapClaims, error) {
	Claims := &jwt.MapClaims{}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// 检查 Token 的签名方法是否与全局配置一致
		if token.Method != SigningMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	}
	token, err := jwt.ParseWithClaims(tokenString, Claims, keyFunc)
	return token.Valid, Claims, err
}
