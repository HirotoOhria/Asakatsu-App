package slack_domain_object

import (
	"strings"

	"github.com/slack-go/slack"
)

type SlackMsgs struct {
	Contents []slack.Message
}

func newSlackMsgs(msgs []slack.Message) *SlackMsgs {
	return &SlackMsgs{
		Contents: msgs,
	}
}

func (d *SlackMsgs) filter(filterStr string) *SlackMsgs {
	var filteredSlackMessages []slack.Message
	for _, msg := range d.Contents {
		if strings.Contains(msg.Text, filterStr) {
			filteredSlackMessages = append(filteredSlackMessages, msg)
		}
	}

	return newSlackMsgs(filteredSlackMessages)
}
