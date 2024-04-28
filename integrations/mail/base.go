package mail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
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

// Mailgun send
func (r *MailRequest) send() bool {
	apiKey := os.Getenv("MAILGUN_API_KEY")
	domain := os.Getenv("MAILGUN_DOMAIN")

	mg := mailgun.NewMailgun(domain, apiKey)

	sender := os.Getenv("EMAIL_FROM")
	subject := r.subject
	body := fmt.Sprintf("From: %s <%s>", r.from, os.Getenv("EMAIL_FROM")) +
		"\r\nTo: " + r.to[0] +
		"\r\nSubject: " +
		r.subject + "\r\n" +
		MIME + "\r\n" +
		r.body
	recipient := r.to[0]

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
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

// net/smtp send
// func (r *MailRequest) send() bool {
// 	apiKey := os.Getenv("MAILGUN_API_KEY")
// 	domain := os.Getenv("MAILGUN_DOMAIN")

// 	// _ = mailgun.NewMailgun(domain, apiKey)
// 	fmt.Println("Mail credentials:", apiKey, domain)

// 	body := fmt.Sprintf("From: %s <%s>", r.from, os.Getenv("EMAIL_FROM")) +
// 		"\r\nTo: " + r.to[0] +
// 		"\r\nSubject: " +
// 		r.subject + "\r\n" +
// 		MIME + "\r\n" +
// 		r.body

// 	SMTP := fmt.Sprintf("%s:%s", os.Getenv("EMAIL_HOST"), os.Getenv("EMAIL_PORT"))

// 	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST")), os.Getenv("EMAIL_FROM"), r.to, []byte(body)); err != nil {
// 		log.Printf("Email sending failed: %s\n", err)
// 		return false
// 	}

// 	// sender := "blessedmadukoma@gmail.com"
// 	// password := apiKey
// 	// recipient := "blessedmadukoma@gmail.com"
// 	// subject := "Fancy subject!"
// 	// body := "Hello from Mailgun Go!"

// 	// // Compose the email
// 	// emailBody := "From: " + sender + "\r\n" +
// 	// 	"To: " + recipient + "\r\n" +
// 	// 	"Subject: " + subject + "\r\n\r\n" +
// 	// 	body

// 	// // Set up authentication information
// 	// auth := smtp.PlainAuth("", sender, password, "smtp.mailgun.org")

// 	// // Connect to the server, authenticate, set the sender and recipient,
// 	// // and send the email body.
// 	// err := smtp.SendMail("smtp.mailgun.org:587", auth, sender, []string{recipient}, []byte(emailBody))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return false
// 	// }

// 	fmt.Println("Email sent successfully!")
// 	return true
// }
