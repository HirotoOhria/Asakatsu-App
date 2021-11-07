package src

import (
	"context"
	"log"
	"net/http"

	"example.com/asakatsu-app/domain/entity/pub_sub_entity"
	"example.com/asakatsu-app/handler"
)

// FetchActivitiesFromSlackBatch は、CloudFunctionsへデプロイする関数
// see https://cloud.google.com/functions/docs/calling/pubsub#sample_code
func FetchActivitiesFromSlackBatch(ctx context.Context, _ pub_sub_entity.PubSubMessage) error {
	if err := handler.FetchActivitiesFromSlackBatch(ctx); err != nil {
		log.Printf("handler.FetchActivitiesFromSlackBatch failed(err=%+v)", err)
		return err
	}

	return nil
}

// GetActivitiesFromSlackUidFunction は、CloudFunctionsへデプロイする関数
// see https://cloud.google.com/functions/docs/calling/http#code_sample
// TODO 関数名をGetActivitiesFromSlackUidApi へ変更し、デプロイする
func GetActivitiesFromSlackUidFunction(w http.ResponseWriter, r *http.Request) {
	_ = handler.GetActivitiesFromSlackUidApi(w, r)
}
