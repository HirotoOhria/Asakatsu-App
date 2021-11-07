package injector

import (
	"context"

	"example.com/asakatsu-app/domain/repository/firestore_repository"
	"example.com/asakatsu-app/domain/repository/slack_repository"
)

func InjectActivityRepostiroy(ctx context.Context) *firestore_repository.ActivityRepository {
	return firestore_repository.NewActivityRepostitory(ctx, InjectFirestoreHandler(ctx))
}

func InjectAsakatsuRepository() *slack_repository.AsakatsuRepository {
	return slack_repository.NewAsakatsuRepository(InjectSlackHandler())
}
