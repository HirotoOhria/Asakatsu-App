package firebase_handler

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
)

type FirebaseHandler struct {
	*firebase.App
}

func NewFirebaseHandlerr(ctx context.Context) *FirebaseHandler {
	app := initFirebaseApp(ctx)

	return &FirebaseHandler{
		App: app,
	}
}

func initFirebaseApp(ctx context.Context) *firebase.App {
	asakatsuAppProjectID := os.Getenv("GCP_PROJECT_ID")

	// Use the application default credentials
	conf := &firebase.Config{
		ProjectID: asakatsuAppProjectID,
	}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		panic(err)
	}

	return app
}
