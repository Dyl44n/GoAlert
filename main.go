package main

import (
	"alertmanager/sms"
)

func main() {
	//	email.SendEmail([]string{"kagarinomiko@gmail.com"},
	//		"Alerta de servidor down",
	//		"Google",
	//		"Erro ao conectar no servidor",
	//		"06/04/2022 15:00",
	//		"./email/template.html")
	// slack.SendSlack("Alerta de servidor down: Google\n Erro: Erro ao conectar no servidor\n Horario: 06/04/2023 15:55\n")
	sms.SendSMS("Alerta de servidor down: Google\n Erro: Erro ao conectar no servidor\n Horario: 06/04/2023 15:55\n", "551199888888")
}
