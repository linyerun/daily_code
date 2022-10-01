package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"redis/config"
)

func Send(subject string, text []byte, toEmails ...string) error {
	em := email.NewEmail()
	em.From = config.EmailFrom
	em.To = toEmails
	em.Subject = subject
	em.Text = text
	return em.Send("smtp.qq.com:25", smtp.PlainAuth("", config.EmailFrom, config.EmailPwd, "smtp.qq.com"))
}

func init() {
	config.InitRedis()
}
