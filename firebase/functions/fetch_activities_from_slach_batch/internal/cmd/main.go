package cmd

import (
	"context"
	"log"
	"time"

	"example.com/fetch-activities-from-slack-batch/internal/firebase/firebase_handler"
	"example.com/fetch-activities-from-slack-batch/internal/firebase/firestore/firestore_handler"
	"example.com/fetch-activities-from-slack-batch/internal/firebase/firestore/firestore_repository"
	"example.com/fetch-activities-from-slack-batch/internal/gcp/pub_sub/pub_sub_entity"
	"example.com/fetch-activities-from-slack-batch/internal/slack/slack_handler"
	"example.com/fetch-activities-from-slack-batch/internal/slack/slack_repository"
)

const location = "Asia/Tokyo"

func init() {
	log.Print("run: cmd.init()")

	initTimeLocation()
}

func initTimeLocation() {
	log.Print("run: cmd.initTimeLocation()")

	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("error: time.LoadLocation(location) err=%+v", err)
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

func FetchActivitiesFromSlackBatch(ctx context.Context, _ pub_sub_entity.PubSubMessage) error {
	log.Print("run: cmd.FetchActivitiesFromSlackBatch()")

	firebaseHandler := firebase_handler.NewFirebaseHandlerr(ctx)

	firestoreHandler := firestore_handler.NewFirestoreHandler(ctx, firebaseHandler)
	defer firestoreHandler.DB.Close()
	activityRepository := firestore_repository.NewActivityRepostitory(ctx, firestoreHandler)

	slackHandler := slack_handler.NewSlackHandler()
	asakatsuRepository := slack_repository.NewAsakatsuRepository(*slackHandler)

	usecase := NewUsecase(asakatsuRepository, activityRepository)
	usecase.Exec()

	return nil
}
