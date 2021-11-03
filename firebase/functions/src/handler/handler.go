package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/asakatsu-app/domain/entity/pub_sub_entity"
	"example.com/asakatsu-app/infrastructure"
	"example.com/asakatsu-app/injector"
)

const location = "Asia/Tokyo"

func init() {
	log.Print("run: handler.init()")

	initTimeLocation()
}

// initTimeLocation は、time パッケージの TimeZone を JST で初期化します。
func initTimeLocation() {
	log.Print("run: handler.initTimeLocation()")

	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("time.LoadLocation fialed.(err=%+v)", err)
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

// FetchActivitiesFromSlackBatch は、FetchActivitiesFromSlackBatch を実行します。
func FetchActivitiesFromSlackBatch(ctx context.Context, _ pub_sub_entity.PubSubMessage) error {
	log.Print("run: handler.fetchActivitiesFromSlackBatch()")

	firebaseHander := injector.InjectFirebaseHandler(ctx)
	firestoreDBConn := infrastructure.GetFirestoreDBConn(ctx, firebaseHander)
	defer firestoreDBConn.Close()

	usecase := injector.InjectFetchActivitiesFromSlackBatchUsecase(ctx)
	if err := usecase.Exec(); err != nil {
		fmt.Printf(err.Error())
	}

	return nil
}
