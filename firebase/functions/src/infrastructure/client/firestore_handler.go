package client

import (
	"context"

	"cloud.google.com/go/firestore"
)

var firestoreHander *FirestoreHandler

type FirestoreHandler struct {
	DB *firestore.Client
}

func NewFirestoreHandler(
	ctx context.Context,
	firebaseHandler *FirebaseHandler,
) *FirestoreHandler {
	if firestoreHander == nil {
		initFirestoreHandler(ctx, firebaseHandler)
	}

	return firestoreHander
}

func GetFirestoreDBConn(ctx context.Context, firebaseHandler *FirebaseHandler) *firestore.Client {
	if firestoreHander == nil {
		initFirestoreHandler(ctx, firebaseHandler)
	}

	return firestoreHander.DB
}

func initFirestoreHandler(ctx context.Context, firebaseHandler *FirebaseHandler) {
	client, err := firebaseHandler.App.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	firestoreHander = &FirestoreHandler{
		DB: client,
	}
}
