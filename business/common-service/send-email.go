package commonservice

import (
	"bytes"
	"change_password_service/api/v1/changepassword/resp"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendEmailNotification(ch resp.DataReset) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading dotenv file")
	}

	var (
		SMTP_HOST    = os.Getenv("SMTP_HOST")
		SMTP_PORT, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
		SMTP_SENDER  = os.Getenv("SMTP_SENDER")
		SMTP_PASS    = os.Getenv("SMTP_PASS")
		SMTP_USER    = os.Getenv("SMTP_USER")
	)

	result, _ := ParseTemplate("templates/email_reset_password.html", ch)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", SMTP_SENDER)
	mailer.SetHeader("To", ch.Email)
	mailer.SetHeader("Subject", "Permintaan reset password")
	mailer.SetBody("text/html", result)

	dialer := gomail.NewDialer(
		SMTP_HOST,
		SMTP_PORT,
		SMTP_USER,
		SMTP_PASS,
	)

	errEmail := dialer.DialAndSend(mailer)
	if errEmail != nil {
		log.Println(errEmail.Error())
		return errEmail
	}

	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	// fmt.Println(err)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func generateId() string {
	id := uuid.New()
	return id.String()
}