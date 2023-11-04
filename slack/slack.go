package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendSlack(textAlerta string) {
	token := os.Getenv("SLACK_AUTH_TOKEN")
	if token == "" {
		panic("SLACK_AUTH_TOKEN não foi configurada")
	}
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		panic("SLACK_CHANNEL_ID não foi configurada")
	}

	client := slack.New(token, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Color:   "danger",
		Pretext: "Alerta de servidor Down",
		Text:    textAlerta,
	}
	_, timestamp, err := client.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Mensagem enviada com sucesso %s as %s", channelID, timestamp)
}
