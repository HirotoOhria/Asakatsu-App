package slack_repository

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"

	"example.com/asakatsu-app/infrastructure"
)

type AsakatsuRepository struct {
	channelID string
	*infrastructure.SlackHandler
}

func NewAsakatsuRepository(slack *infrastructure.SlackHandler) *AsakatsuRepository {
	channelID := os.Getenv("SLACK_ASAKATSU_CHANNEL_ID")

	return &AsakatsuRepository{
		channelID:    channelID,
		SlackHandler: slack,
	}
}

// GetYesterdayConversationHistory は、昨日の会話を取得します。
// 例えば、9月1日に実行した場合、8月31日の00時00分~23時59分までの会話を取得します。
// see https://api.slack.com/methods/conversations.history
// see https://pkg.go.dev/github.com/slack-go/slack#Client.GetConversationHistory
func (r *AsakatsuRepository) GetYesterdayConversationHistory() ([]slack.Message, error) {
	yesterday := time.Now().Add(-24 * time.Hour)
	startFetchPeriodDate := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local)
	endFetchPeriodDate := startFetchPeriodDate.Add(24 * time.Hour)

	params := &slack.GetConversationHistoryParameters{
		ChannelID: r.channelID,
		Oldest:    strconv.FormatInt(startFetchPeriodDate.Unix(), 10),
		Latest:    strconv.FormatInt(endFetchPeriodDate.Unix(), 10),
	}

	response, err := r.Api.GetConversationHistory(params)
	if err != nil {
		log.Printf("slack.Client.GetConversationHistory failed(err=%+v)", err)
		return nil, err
	}

	return response.Messages, nil
}

// GetConversationReplies は、会話のリプライメッセージを取得します。
// see https://api.slack.com/methods/conversations.replies
// see https://pkg.go.dev/github.com/nlopes/slack#Client.GetConversationReplies
func (r *AsakatsuRepository) GetConversationReplies(timestamp string) ([]slack.Message, error) {
	params := &slack.GetConversationRepliesParameters{
		ChannelID: r.channelID,
		Timestamp: timestamp,
	}

	msgs, _, _, err := r.Api.GetConversationReplies(params)
	if err != nil {
		log.Printf("slack.Client.GetConversationReplies failed(err=%+v)", err)
		return nil, err
	}

	return msgs, err
}
