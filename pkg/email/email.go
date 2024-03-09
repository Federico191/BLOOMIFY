package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"projectIntern/internal/entity"
	"projectIntern/pkg/config"
	"strconv"
)

const appName = "Bloomify"

type EmailItf interface {
	SendEmailVerification(user *entity.User, verificationCode string) error
}

type Email struct {
	env *config.Env
}

func NewEmail(env *config.Env) EmailItf {
	return &Email{env: env}
}

func (e Email) SendEmailVerification(user *entity.User, verificationCode string) error {
	url := "http://" + e.env.AHost + e.env.APort + "/" + "verify_email" + "/" + verificationCode

	textString := fmt.Sprintf(
		"Dear %s,\n\n"+
			"Thank you for registering with %s. To complete the registration process, you need to verify your email.\n\n"+
			"Please click the link below:\n\n%s\n\n"+
			"If you did not request registration with %s, you can ignore this email.\n\n"+
			"Thank you.\n\n"+
			"Regards,\n"+
			"The %s Team\n",
		user.FullName,
		appName,
		url,
		appName,
		appName,
	)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.env.EmailFrom)
	mailer.SetHeader("To", user.Email)
	mailer.SetHeader("Subject", "Your email verification")
	mailer.SetBody("text/html", textString)

	port, err := strconv.Atoi(e.env.SMTPPort)
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(e.env.SMTPHost, port, e.env.SMTPUser, e.env.SMTPPassword)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
