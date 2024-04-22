package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type MailRequest struct {
	from    string
	to      []string
	subject string
	body    string
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func newMailRequest(to []string, subject string) *MailRequest {
	return &MailRequest{
		from:    os.Getenv("APP_NAME"),
		to:      to,
		subject: subject,
	}
}

func (r *MailRequest) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *MailRequest) send() bool {
	body := fmt.Sprintf("From: %s <%s>", r.from, os.Getenv("EMAIL_FROM")) +
		"\r\nTo: " + r.to[0] +
		"\r\nSubject: " +
		r.subject + "\r\n" +
		MIME + "\r\n" +
		r.body

	SMTP := fmt.Sprintf("%s:%s", os.Getenv("EMAIL_HOST"), os.Getenv("EMAIL_PORT"))

	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST")), os.Getenv("EMAIL_FROM"), r.to, []byte(body)); err != nil {
		log.Printf("Email sending failed: %s\n", err)
		return false
	}
	return true
}

func (r *MailRequest) SendMail(templateName string, items interface{}) error {

	err := r.parseTemplate(templateName, items)
	if err != nil {
		return err
	}
	if ok := r.send(); ok {
		log.Printf("Email has been sent to %s\n", r.to)
		return nil
	}

	return fmt.Errorf("failed to send the email to %s", r.to)

}
