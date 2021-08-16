package firestore_handler

import (
	"context"

	"example.com/fetch-activities-from-slack-batch/internal/firebase/firebase_handler"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreHandler struct {
	DB *firestore.Client
}

func NewFirestoreHandler(
	ctx context.Context,
	firebaseHandler *firebase_handler.FirebaseHandler,
) *FirestoreHandler {
	client := initFirestore(firebaseHandler.App, ctx)

	return &FirestoreHandler{
		DB: client,
	}
}

func initFirestore(firebaseApp *firebase.App, ctx context.Context) *firestore.Client {
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	return client
}
