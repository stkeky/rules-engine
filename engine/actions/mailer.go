package actions

import (
	"net/smtp"
	"fmt"
	"sync"
	"net/mail"
	"crypto/tls"
	"os"
)

const (
	host             string = "smtp.gmail.com"
	port             int    = 465
	envEmailUsername string = "EMAIL_USERNAME"
	envEmailPassword string = "EMAIL_PASSWORD"
	fromHeader       string = "From"
	toHeader         string = "To"
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
func (ms *mailService) SendEmail(recipient, content string) {
	from := mail.Address{Address: ms.username}
	to := mail.Address{Address: recipient}

	headers := make(map[string]string)
	headers[fromHeader] = from.String()
	headers[toHeader] = to.String()

	var message string
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + content

	serverName := fmt.Sprintf("%s:%d", host, port)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", serverName, tlsConfig)
	if err != nil {
		fmt.Println(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		fmt.Println(err)
	}

	if err = c.Auth(ms.auth); err != nil {
		fmt.Println(err)
	}

	if err = c.Mail(from.Address); err != nil {
		fmt.Println(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		fmt.Println(err)
	}

	w, err := c.Data()
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
	}

	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}

	c.Quit()
}
