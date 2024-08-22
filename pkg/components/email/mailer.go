package email

import (
	"crypto/tls"

	"github.com/fimreal/goutils/ezap"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// Letter represents the content of an email.
type Letter struct {
	Sender    string `json:"sender" form:"sender" validate:"required"`       // Email sender
	Recipient string `json:"recipient" form:"recipient" validate:"required"` // Email recipient
	Subject   string `json:"subject" form:"subject" validate:"required"`     // Email subject
	Body      string `json:"body" form:"body" validate:"required"`           // Email body
	Type      string `json:"type" form:"type" validate:"required"`           // Email content type (text/html, text/plain, etc.)
}

// Mailer is responsible for sending emails.
type Mailer struct {
	Username           string `validate:"email"`
	Password           string
	SmtpHost           string
	SmtpPort           int
	InsecureSkipVerify bool
}

// ViperMailer retrieves mailer configuration from viper.
func ViperMailer() *Mailer {
	return &Mailer{
		Username:           viper.GetString("email_username"),
		Password:           viper.GetString("email_password"),
		SmtpHost:           viper.GetString("email_smtp_host"),
		SmtpPort:           viper.GetInt("email_smtp_port"),
		InsecureSkipVerify: viper.GetBool("email_Insecureskipverify"),
	}
}

// NewMailer creates a new Mailer instance.
func NewMailer(username, password, smtpHost string, smtpPort int, InsecureSkipVerify bool) *Mailer {
	return &Mailer{
		Username:           username,
		Password:           password,
		SmtpHost:           smtpHost,
		SmtpPort:           smtpPort,
		InsecureSkipVerify: InsecureSkipVerify,
	}
}

// Send sends an email using the Mailer instance.
func (mailer *Mailer) Send(letter *Letter) error {
	host := mailer.SmtpHost
	port := mailer.SmtpPort
	username := mailer.Username
	password := mailer.Password
	ezap.Debugf("mailer: smtp host: %s:%d, username: %s", host, port, username)

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(username, letter.Sender))
	m.SetHeader("To", letter.Recipient)
	m.SetHeader("Subject", letter.Subject)
	m.SetBody(letter.Type, letter.Body)

	d := gomail.NewDialer(host, port, username, password)
	ezap.Debugf("letter: mailto: %s, subject: %s, type: %s, body: %s", letter.Recipient, letter.Subject, letter.Type, letter.Body)

	// Skip TLS authentication to solve the error problem of "x509: certificate signed by unknown authority" due to the lack of certificate files in the system.
	if mailer.InsecureSkipVerify {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		ezap.Debug("mailer: skip TLS authentication")
	}

	ezap.Debugf("mailer: dial and send")
	return d.DialAndSend(m)
}
