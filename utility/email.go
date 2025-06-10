package utility

import (
	"investment/server/conf"

	"gopkg.in/gomail.v2"
)

func SendHtmlMail(subject, body string) error {
	host := conf.Get().System.MailServer
	port := conf.Get().System.MailServerPort
	user := conf.Get().System.FromMail
	pw := conf.Get().System.FromMailPassword

	msg := gomail.NewMessage()
	msg.SetHeader("From", "investment ai"+"<"+user+">")
	msg.SetHeader("To", conf.Get().System.ToMailList...)
	if len(conf.Get().System.ToMailList) > 0 {
		msg.SetHeader("Cc", conf.Get().System.ToMailList...)
	}
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	return gomail.NewDialer(host, port, user, pw).DialAndSend(msg)
}
