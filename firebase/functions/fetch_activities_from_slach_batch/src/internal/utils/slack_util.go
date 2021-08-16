package utils

import (
	"strings"

	"github.com/slack-go/slack"
)

func FindStartConversations(targetConversations []slack.Message) []slack.Message {
	return filterConversations(targetConversations, ":start:")
}

func FindEndConversations(targetConversations []slack.Message) []slack.Message {
	return filterConversations(targetConversations, ":end1:")
}

func filterConversations(targetConversations []slack.Message, filterStr string) []slack.Message {
	var filteredConversations []slack.Message

	for _, targetConversation := range targetConversations {
		if strings.Contains(targetConversation.Text, filterStr) {
			filteredConversations = append(filteredConversations, targetConversation)
		}
	}

	return filteredConversations
}
