package slack_domain_object

import (
	"github.com/slack-go/slack"
)

type SlackReplies struct {
	SlackMsgs
}

func NewSlackReplies(msgs []slack.Message) *SlackReplies {
	return &SlackReplies{
		SlackMsgs: *newSlackMsgs(msgs),
	}
}

func (d *SlackReplies) FindEndMsg() *SlackEndMsg {
	endSlackMsgs := d.findEndSlackMsgs()
	if len(endSlackMsgs.Contents) == 0 {
		return nil
	}

	return NewSlackEndMsg(endSlackMsgs.Contents[0])
}

func (d *SlackReplies) findEndSlackMsgs() *SlackMsgs {
	return d.filter(":end1:")
}
