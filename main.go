package main

import (
	"alertmanager/email"
)

func main() {
	email.SendEmail([]string{"kagarinomiko@gmail.com"},
		"Alerta de servidor down",
		"Google",
		"Erro ao conectar no servidor",
		"06/04/2022 15:00",
		"./email/template.html")
}
