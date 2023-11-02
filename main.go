package main

import (
	"net/smtp"
	"os"
	"text/template"
)

func main() {
	from := "kagarinomiko@gmail.com"
	password := os.Getenv("GMAIL_PASSWORD")
	if password == "" {
		panic("GMAIL_PASSWORD não foi configurado")
		os.Exit(1)
	}
	to := []string{
		"kagarinomiko@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	t, _ := template.ParseFiles("template.html")
}
