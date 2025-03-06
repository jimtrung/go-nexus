package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HasMXRecord(email string) bool {
	domain := email[strings.LastIndex(email, "@")+1:]
	mxRecords, err := net.LookupMX(domain)
	return err == nil && len(mxRecords) > 0
}

func GenerateToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func SendVerificationEmail(email, token string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := []byte(fmt.Sprintf("Subject: Verify Your Email\n\nClick here to verify: http://127.0.0.1:8080/p/user/verify/%s",
        token,
    ))

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
}

func ResetPasswordEmail(email, token string) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := []byte(fmt.Sprintf("Subject: Reset Password\n\nClick here to reset your password: http://127.0.0.1:8080/p/user/reset-password/%s",
        token,
    ))

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
}
