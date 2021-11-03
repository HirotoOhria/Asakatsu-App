package injector

import (
	"example.com/asakatsu-app/domain/repository/slack_repository"
	"example.com/asakatsu-app/infrastructure"
)

func InjectSlackHandler() *infrastructure.SlackHandler {
	return infrastructure.NewSlackHandler()
}

func InjectAsakatsuRepository() *slack_repository.AsakatsuRepository {
	return slack_repository.NewAsakatsuRepository(InjectSlackHandler())
}
