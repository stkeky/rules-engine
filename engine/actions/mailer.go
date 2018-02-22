package actions

import (
	"net/smtp"
	"fmt"
	"sync"
	"os"
)

const (
	host             string = "smtp.gmail.com"
	port             int    = 587
	envEmailUsername string = "EMAIL_USERNAME"
	envEmailPassword string = "EMAIL_PASSWORD"
)

type mailService struct {
	auth     smtp.Auth
	username string
}

var instance *mailService
var once sync.Once

// MailService return singleton mailing service
func MailService() *mailService {
	once.Do(func() {
		username := os.Getenv(envEmailUsername)

		auth := smtp.PlainAuth(
			"",
			username,
			os.Getenv(envEmailPassword),
			host)

		instance = &mailService{auth, username}
	})
	return instance
}

// SendEmail sends email to specified address with specified content
func (ms *mailService) SendEmail(to, content string) {
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", host, port),
		ms.auth,
		ms.username,
		[]string{to},
		[]byte(content),
	)

	if err != nil {
		fmt.Println(err)
	}
}
