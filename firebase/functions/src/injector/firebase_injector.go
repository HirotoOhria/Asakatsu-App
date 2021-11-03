package injector

import (
	"context"

	"example.com/asakatsu-app/infrastructure"
)

func InjectFirebaseHandler(ctx context.Context) *infrastructure.FirebaseHandler {
	return infrastructure.NewFirebaseHandlerr(ctx)
}
