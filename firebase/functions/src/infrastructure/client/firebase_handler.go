package client

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
)

type FirebaseHandler struct {
	*firebase.App
}

func NewFirebaseHandler(ctx context.Context) *FirebaseHandler {
	asakatsuAppProjectID := os.Getenv("GCP_PROJECT_ID")

	// Use the application default credentials
	conf := &firebase.Config{
		ProjectID: asakatsuAppProjectID,
	}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		panic(err)
	}

	return &FirebaseHandler{
		App: app,
	}
}
