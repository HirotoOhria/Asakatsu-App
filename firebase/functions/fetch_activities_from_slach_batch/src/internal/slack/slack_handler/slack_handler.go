package slack_handler

import (
	"os"

	"github.com/slack-go/slack"
)

type SlackHandler struct {
	Api *slack.Client
}

func NewSlackHandler() *SlackHandler {
	slackApi := initSlackApi()

	return &SlackHandler{
		Api: slackApi,
	}
}

func initSlackApi() *slack.Client {
	slackBotOAuthToken := os.Getenv("SLACK_BOT_OAUTH_TOKEN")

	return slack.New(slackBotOAuthToken, slack.OptionDebug(true))
}
