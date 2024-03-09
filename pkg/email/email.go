package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"projectIntern/internal/entity"
	"projectIntern/pkg/config"
)

const appName = "Bloomify"

type EmailItf interface {
	SendEmailVerification(user *entity.User, verificationUrl string) error
}

type Email struct {
	email *email.Email
	env   *config.Env
}

func NewEmail(email *email.Email, env *config.Env) EmailItf {
	return &Email{email: email, env: env}
}

func (e Email) SendEmailVerification(user *entity.User, verificationUrl string) error {
	auth := smtp.PlainAuth("", e.env.EmailFrom, e.env.SMTPPassword, e.env.SMTPHost)

	textString := fmt.Sprintf("Dear %s,\n\nThank you for registering with %s. To complete the registration process, you need to verify your email.\n\nPlease click the link below:\n\n%s\n\nIf you did not request registration with %s, you can ignore this email.\n\nThank you.\n\nRegards,\nThe %s Team\n", user.FullName, appName, verificationUrl, appName, appName)

	text := []byte(textString)

	createdEmail := email.Email{
		From:    e.env.EmailFrom,
		To:      []string{user.Email},
		Subject: "Your account verification",
		Text:    text,
	}

	return createdEmail.Send(user.Email, auth)
}
