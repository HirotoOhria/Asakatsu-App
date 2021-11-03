package src

import (
	"context"
	"log"

	"example.com/asakatsu-app/domain/entity/pub_sub_entity"
	"example.com/asakatsu-app/handler"
)

// FetchActivitiesFromSlackBatch is function for deploy.
// see https://cloud.google.com/functions/docs/calling/pubsub?hl=ja#sample_code
func FetchActivitiesFromSlackBatch(ctx context.Context, m pub_sub_entity.PubSubMessage) error {
	if err := handler.FetchActivitiesFromSlackBatch(ctx, m); err != nil {
		log.Printf("handler.FetchActivitiesFromSlackBatch failed(err=%+v)", err)
		return err
	}

	return nil
}
