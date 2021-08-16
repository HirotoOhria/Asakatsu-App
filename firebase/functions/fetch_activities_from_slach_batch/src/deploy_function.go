package src

import (
	"context"
	"log"

	"example.com/fetch-activities-from-slack-batch/cmd/cloud_funcsions"
	"example.com/fetch-activities-from-slack-batch/internal/gcp/pub_sub/pub_sub_entity"
)

// FetchActivitiesFromSlackBatch is function for deploy.
// see https://cloud.google.com/functions/docs/calling/pubsub?hl=ja#sample_code
func FetchActivitiesFromSlackBatch(ctx context.Context, m pub_sub_entity.PubSubMessage) error {
	if err := cloud_funcsions.Main(ctx, m); err != nil {
		log.Printf("error: cmd.FetchActivitiesFromSlackBatch() err=%+v", err)
		return err
	}

	return nil
}
