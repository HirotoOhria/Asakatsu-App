package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/asakatsu-app/domain/api_io"

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
func GetActivitiesFromSlackUidFunction(w http.ResponseWriter, _ *http.Request) {
	input := &api_io.GetActivitiesFromSlackUidInput{
		SlackUID: "U01DG785DT4",
	}

	output, err := handler.GetActivitiesFromSlackUidFunction(input)
	if err != nil {
		log.Fatalf("handler.GetActivitiesFromSlackUidFunction failed(err=%+v)", err)
	}
	if output == nil {
		log.Fatal("handler.GetActivitiesFromSlackUidFunction failed")
	}

	w.Header().Set("Content-Type", "application/json")

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatal("json.Marshal failed")
	}

	fmt.Fprintf(w, string(outputJson))
}
