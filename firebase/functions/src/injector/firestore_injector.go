package injector

import (
	"context"

	"example.com/asakatsu-app/domain/repository/firestore_repository"
	"example.com/asakatsu-app/infrastructure"
)

func InjectFirestoreHandler(ctx context.Context) *infrastructure.FirestoreHandler {
	return infrastructure.NewFirestoreHandler(ctx, InjectFirebaseHandler(ctx))
}

func InjectActivityRepostiroy(ctx context.Context) *firestore_repository.ActivityRepository {
	return firestore_repository.NewActivityRepostitory(ctx, InjectFirestoreHandler(ctx))
}
