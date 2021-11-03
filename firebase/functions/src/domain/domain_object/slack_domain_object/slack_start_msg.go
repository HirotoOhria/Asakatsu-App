package slack_domain_object

import (
	"time"

	"example.com/asakatsu-app/util"
	"github.com/slack-go/slack"
)

type SlackStartMsg struct {
	Content slack.Message
}

func NewSlackStartMsg(msg slack.Message) *SlackStartMsg {
	return &SlackStartMsg{
		Content: msg,
	}
}

func (d *SlackStartMsg) GetTime() *time.Time {
	t, _ := util.ParseTimeFromFloatStr(d.Content.Timestamp)
	return t
}
