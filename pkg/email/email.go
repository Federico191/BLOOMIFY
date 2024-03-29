package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"projectIntern/internal/entity"
	"strconv"
)

const appName = "Bloomify"

type EmailItf interface {
	SendEmailVerification(user *entity.User, verificationCode string) error
}

type Email struct {
}

func NewEmail() EmailItf {
	return &Email{}
}

func (e Email) SendEmailVerification(user *entity.User, verificationCode string) error {
	url := "https://bloomify2-cc614ecf7708.herokuapp.com" + "/api/v1/user/" + "verify-email" + "/" + verificationCode

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
	mailer.SetHeader("From", "bloomify20@gmail.com")
	mailer.SetHeader("To", fmt.Sprintf("<%s>", user.Email))
	mailer.SetHeader("Subject", "Your email verification")
	mailer.SetBody("text/html", textString)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
