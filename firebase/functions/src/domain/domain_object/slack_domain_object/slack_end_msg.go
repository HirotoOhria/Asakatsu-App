package slack_domain_object

import (
	"time"

	"example.com/asakatsu-app/util"
	"github.com/slack-go/slack"
)

type SlackEndMsg struct {
	Content slack.Message
}

func NewSlackEndMsg(msg slack.Message) *SlackEndMsg {
	return &SlackEndMsg{
		Content: msg,
	}
}

func (d *SlackEndMsg) GetTime() *time.Time {
	t, _ := util.ParseTimeFromFloatStr(d.Content.Timestamp)
	return t
}
