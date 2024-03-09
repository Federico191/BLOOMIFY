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
	url := "http://" + e.env.AHost + e.env.APort + "/api/v1" + "/" + "verify_email" + "/" + verificationCode

	textString := fmt.Sprintf(`
    <html>
    <head>
        <style>
            body {
                font-family: Arial, sans-serif;
            }
            .container {
                max-width: 600px;
                margin: 0 auto;
            }
            .header {
                background-color: #f2f2f2;
                padding: 20px;
                text-align: center;
            }
            .content {
                padding: 20px;
            }
            .button {
                display: inline-block;
                background-color: #007bff;
                color: #fff;
                padding: 10px 20px;
                text-decoration: none;
                border-radius: 5px;
            }
        </style>
    </head>
    <body>
        <div class="container">
            <div class="header">
                <h2>Thank You for Registering with %s</h2>
            </div>
            <div class="content">
                <p>Dear %s,</p>
                <p>Thank you for registering with %s. To complete the registration process, you need to verify your email.</p>
                <p>Please click the button below to verify your email:</p>
                <a href="%s" class="button">Verify Email</a>
                <p>If you did not request registration with %s, you can ignore this email.</p>
                <p>Thank you.</p>
                <p>Regards,<br/>The %s Team</p>
            </div>
        </div>
    </body>
    </html>
`,
		appName, user.FullName, appName, url, appName, appName)

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
