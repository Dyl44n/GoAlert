package main

import (
	"bytes"
	"fmt"
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

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Alerta: Servidor down \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Server  string
		Error   string
		Horario string
	}{
		Server:  "Google",
		Error:   "Erro ao acessar o servidor.",
		Horario: "10/12/2022 14:00",
	})
}
