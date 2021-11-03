package infrastructure

import (
	"os"

	"github.com/slack-go/slack"
)

type SlackHandler struct {
	Api *slack.Client
}

func NewSlackHandler() *SlackHandler {
	slackBotOAuthToken := os.Getenv("SLACK_BOT_OAUTH_TOKEN")
	slackApi := slack.New(slackBotOAuthToken, slack.OptionDebug(true))

	return &SlackHandler{
		Api: slackApi,
	}
}
