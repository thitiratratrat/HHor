package service

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type EmailService interface {
	SendEmail(recipient string, code string)
}

func EmailServiceHandler() EmailService {
	return &emailService{}
}

type emailService struct{}

func (emailService *emailService) SendEmail(recipient string, code string) {
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")

	to := []string{
		recipient,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	template, templateErr := template.ParseFiles("./template/template.html")

	if templateErr != nil {
		panic(templateErr)
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: HHOR Verification Code \n%s\n\n", mimeHeaders)))

	template.Execute(&body, struct {
		Code string
	}{
		Code: code,
	})

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

	if err != nil {
		panic(err)
	}
}
