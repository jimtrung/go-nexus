package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func IsValidEmail(email string) bool {
	return true
}

func HasMXRecord(email string) bool {
	domain := email[strings.LastIndex(email, "@")+1:]
	mxRecords, err := net.LookupMX(domain)
	return err == nil && len(mxRecords) > 0
}

func GenerateToken() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println("Error generating token:", err)
		return ""
	}
	return hex.EncodeToString(bytes)
}

func SendEmail(to, subject, body string) error {
	fromEmail := os.Getenv("EMAIL")
	apiKey := os.Getenv("SENDGRID_API_KEY")

	if fromEmail == "" || apiKey == "" {
		log.Println("Missing EMAIL or SENDGRID_API_KEY environment variables")
		return fmt.Errorf("missing email credentials")
	}

	from := mail.NewEmail("GoNexus", fromEmail)
	toEmail := mail.NewEmail("", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, body)

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Printf("Email sent to %s. Status Code: %d\n", to, response.StatusCode)
	return nil
}

func SendVerificationEmail(email, token string) {
	body := fmt.Sprintf("Click here to verify: http://127.0.0.1:8080/auth/verify/%s", token)
	go SendEmail(email, "Verify Your Email", body)
}

func ResetPasswordEmail(email, token string) {
	body := fmt.Sprintf("Click here to reset password: http://127.0.0.1:8080/reset-password/%s", token)
	go SendEmail(email, "Reset Password", body)
}
