package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	// Set up authentication information.
	smtpHost := os.Getenv("smtpEmailHost") // Replace with your SMTP server
	smtpPort := os.Getenv("smtpEmailPort") // Replace with your SMTP port (587 or 465 for SSL)
	smtpUser := os.Getenv("smtpUserEmail") // Replace with your email
	smtpPass := os.Getenv("smtpuserPass")  // Replace with your email password

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Create the email headers and body.
	from := smtpUser
	To := []string{to}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body + "\r\n")

	// Send the email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, To, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
