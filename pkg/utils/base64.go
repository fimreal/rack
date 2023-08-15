package utils

import (
	"encoding/base64"
)

func Base64Encode(s string) string {
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	return s64
}

func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// 自定义 base64 转换内容, 例如 encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
func Base64Encodef(s string, encodeStd string) string {
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	return s64
}

func Base64UrlEncode(s string) string {
	s64 := base64.URLEncoding.EncodeToString([]byte(s))
	return s64
}

func BaseUrl64Decode(s string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
