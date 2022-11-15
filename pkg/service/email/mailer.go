package email

import (
	"crypto/tls"
	"os"
	"strconv"

	"github.com/fimreal/goutils/ezap"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Username       string `validate:"email"`
	Password       string
	SmtpServer     string
	SmtpServerPort string
}

func NewMailer() *Mailer {
	// return &Mailer{
	// 	Username:       viper.GetString("username"),
	// 	Password:       viper.GetString("password"),
	// 	Smtpserver:     viper.GetString("smtpserver"),
	// 	Smtpserverport: viper.GetString("smtpserverport"),
	// }
	return &Mailer{
		Username:       os.Getenv("MAILUSERNAME"),
		Password:       os.Getenv("MAILPASSWORD"),
		SmtpServer:     os.Getenv("MAILSMTPSERVER"),
		SmtpServerPort: os.Getenv("MAILSMTPSERVERPORT"),
	}
}

// Mailto 发送电子邮件
func Mailto(letter *Letter) error {
	mailer := NewMailer()
	host := mailer.SmtpServer
	port, _ := strconv.Atoi(mailer.SmtpServerPort)
	username := mailer.Username
	password := mailer.Password
	ezap.Debugf("邮箱配置 smtp 服务器: %s:%d, 用户名: %s, 密码: ***", host, port, username)

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(username, "GoMail Robot"))
	m.SetHeader("To", letter.Mailto...)
	m.SetHeader("Subject", letter.Subject)
	m.SetBody("text/html", letter.Body)

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 解决 x509: certificate signed by unknown authority 报错问题, 关掉 tls 认证

	return d.DialAndSend(m)
}
