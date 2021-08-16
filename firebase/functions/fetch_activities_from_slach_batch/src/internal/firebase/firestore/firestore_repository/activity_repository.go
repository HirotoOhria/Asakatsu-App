package firestore_repository

import (
	"context"
	"log"

	"example.com/fetch-activities-from-slack-batch/internal/firebase/firestore/firestore_entity"
	"example.com/fetch-activities-from-slack-batch/internal/firebase/firestore/firestore_handler"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const collectionPath = "activities"

type ActivityRepository struct {
	ctx context.Context
	db  *firestore.Client
}

func NewActivityRepostitory(
	ctx context.Context,
	firestoreHandler *firestore_handler.FirestoreHandler,
) *ActivityRepository {
	return &ActivityRepository{
		ctx: ctx,
		db:  firestoreHandler.DB,
	}
}

func (r *ActivityRepository) Get(docID string) (*firestore_entity.ActivityDoc, error) {
	dataSnap, err := r.db.Collection(collectionPath).Doc(docID).Get(r.ctx)
	if status.Code(err) == codes.NotFound {
		log.Fatalf("docment is not found(err=%+v)", err)
		return nil, err
	} else if err != nil {
		log.Fatalf("get dockment failed(err=%+v)", err)
		return nil, err
	}

	activityDoc := &firestore_entity.ActivityDoc{
		ID: docID,
	}

	err = dataSnap.DataTo(&activityDoc.Data)
	if err != nil {
		log.Fatalf("document transform to entity failed(err=%+v)", err)
		return nil, err
	}

	return activityDoc, nil
}

func (r *ActivityRepository) Set(activityDoc firestore_entity.ActivityDoc) error {
	_, err := r.db.Collection(collectionPath).Doc(activityDoc.ID).Set(r.ctx, activityDoc.Data)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return err
	}

	return nil
}
