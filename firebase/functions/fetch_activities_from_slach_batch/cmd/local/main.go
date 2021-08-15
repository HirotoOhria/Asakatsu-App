package main

import (
	"context"
	"log"

	"example.com/fetch-activities-from-slack-batch/internal/cmd"
	"example.com/fetch-activities-from-slack-batch/internal/gcp/pub_sub/pub_sub_entity"
	"github.com/joho/godotenv"
)

func init() {
	log.Print("run: main.init()")

	initDotenv()
}

func initDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error: loading .env file")
	}
}

func main() {
	log.Print("run: main.main()")

	ctx := context.Background()
	msg := new(pub_sub_entity.PubSubMessage)

	if err := cmd.FetchActivitiesFromSlackBatch(ctx, *msg); err != nil {
		log.Fatalf("error: FetchActivitiesFromSlackBatch() err=%+v", err)
	}
}
