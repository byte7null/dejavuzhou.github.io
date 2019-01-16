package robot

import (
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
	"strings"
)

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func sendMail(user, password, host, to, subject, body, mailtype string) error {
	auth := unencryptedAuth{
		smtp.PlainAuth(
			"",
			user,
			password,
			host,
		),
	}
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	viper.GetString("")
	err := smtp.SendMail(viper.GetString("STMP_HOST")+":"+viper.GetString("STMP_PORT"), auth, user, send_to, msg)
	return err
}

func SendMsgToEmail(subject, msg string) error {
	err := sendMail(viper.GetString("STMP_USER"), viper.GetString("STMP_PASSWORD"), viper.GetString("STMP_HOST"), viper.GetString("STMP_RECEIVER"), subject, msg, "html")
	if err != nil {
		fmt.Println(err)
	}
	return err
}
