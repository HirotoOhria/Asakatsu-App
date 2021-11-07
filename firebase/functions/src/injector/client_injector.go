package injector

import (
	"context"

	"example.com/asakatsu-app/infrastructure/client"
)

// TODO SDK 設定の構造体名を、HandlerからClientへ修正する
func InjectFirebaseHandler(ctx context.Context) *client.FirebaseHandler {
	return client.NewFirebaseHandler(ctx)
}

func InjectFirestoreHandler(ctx context.Context) *client.FirestoreHandler {
	return client.NewFirestoreHandler(ctx, InjectFirebaseHandler(ctx))
}

func InjectSlackHandler() *client.SlackHandler {
	return client.NewSlackHandler()
}
