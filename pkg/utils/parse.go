package utils

import (
	"net"
	"regexp"
	"strings"
)

func IsIPv4(ip string) bool {
	ipaddr := net.ParseIP(ip)

	if ipaddr == nil {
		return false
	}

	v4 := ipaddr.To4()
	return v4 != nil
}

func IsLanIPv4(ip string) bool {
	if !IsIPv4(ip) {
		return false
	}

	lanex := []string{"10.", "192.168.", "172.16.", "172.17.", "172.18.", "172.19.", "172.20.", "172.21.", "172.22.", "172.23.", "172.24.", "172.25.", "172.26.", "172.27.", "172.28.", "172.29.", "172.30.", "172.31.", "169.254."}
	for _, lan := range lanex {
		if strings.HasPrefix(ip, lan) {
			return true
		}
	}
	return false
}

func IsEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{2,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// 国内手机号
func IsPhoneNumber(phone string) bool {
	pattern := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}

func IsUsername(username string) bool {
	pattern := `^[a-z][_.-a-z0-9]{2,31}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(username)
}
