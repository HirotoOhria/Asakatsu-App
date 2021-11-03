package slack_domain_object

import (
	"github.com/slack-go/slack"
)

type SlackConversation struct {
	SlackMsgs
}

func NewSlackConversation(msgs []slack.Message) *SlackConversation {
	return &SlackConversation{
		SlackMsgs: *newSlackMsgs(msgs),
	}
}

func (d *SlackConversation) FindStartSlackMsgs() []SlackStartMsg {
	startSlackMsgs := d.findStartSlackMsgs()
	if len(startSlackMsgs.Contents) == 0 {
		return nil
	}

	var slackStartMsgs []SlackStartMsg
	for _, slackMsg := range startSlackMsgs.Contents {
		slackStartMsgs = append(slackStartMsgs, *NewSlackStartMsg(slackMsg))
	}

	return slackStartMsgs
}

func (d *SlackConversation) findStartSlackMsgs() *SlackMsgs {
	return d.filter(":start:")
}
