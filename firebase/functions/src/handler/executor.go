package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"example.com/asakatsu-app/domain/api_io"
	"example.com/asakatsu-app/infrastructure/client"
	"example.com/asakatsu-app/injector"
)

const location = "Asia/Tokyo"

func init() {
	log.Print("run: executor.init()")

	initTimeLocation()
}

// initTimeLocation は、time パッケージの TimeZone を JST で初期化します。
func initTimeLocation() {
	log.Print("run: executor.initTimeLocation()")

	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Printf("time.LoadLocation fialed.(err=%+v)", err)
		loc = time.FixedZone(location, 9*60*60)
	}

	time.Local = loc
}

func FetchActivitiesFromSlackBatch(ctx context.Context) error {
	log.Print("run: executor.FetchActivitiesFromSlackBatch()")

	firebaseHander := injector.InjectFirebaseHandler(ctx)
	firestoreDBConn := client.GetFirestoreDBConn(ctx, firebaseHander)
	defer firestoreDBConn.Close()

	handler := injector.InjectFetchActivitiesFromSlackHandler(ctx)
	if err := handler.Exec(); err != nil {
		log.Printf("FetchActivitiesFromSlackHandler.Exec failed(err=%+v)", err)
		return err
	}

	return nil
}

func GetActivitiesFromSlackUidApi(
	w http.ResponseWriter, r *http.Request,
) *api_io.GetActivitiesFromSlackUidOutput {
	log.Print("run: executor.GetActivitiesFromSlackUidOutput()")

	ctx := context.Background()
	firebaseHander := injector.InjectFirebaseHandler(ctx)
	firestoreDBConn := client.GetFirestoreDBConn(ctx, firebaseHander)
	defer firestoreDBConn.Close()

	handler := injector.InjectGetActivitiesFromSlackUidHandler(ctx)
	output := handler.Exec(w, r)

	return output
}
