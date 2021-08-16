package cloud_funcsions

import (
	"context"
	"log"

	"example.com/fetch-activities-from-slack-batch/internal/cmd"
	"example.com/fetch-activities-from-slack-batch/internal/gcp/pub_sub/pub_sub_entity"
)

func Main(ctx context.Context, m pub_sub_entity.PubSubMessage) error {
	err := cmd.FetchActivitiesFromSlackBatch(ctx, m)
	if err != nil {
		log.Fatalf("error: FetchActivitiesFromSlackBatch() err=%+v", err)
		return err
	}

	return nil
}
