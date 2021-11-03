package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"

	"example.com/asakatsu-app/domain/entity/pub_sub_entity"
	"example.com/asakatsu-app/handler"
)

func init() {
	log.Print("run: main.init()")

	initDotenv()
}

// initDotenv は、 .env ファイルの環境変数を読み込みます。
func initDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error: loading .env file")
	}
}

func main() {
	log.Print("run: main.main()")

	ctx := context.Background()
	msg := new(pub_sub_entity.PubSubMessage)

	if err := handler.FetchActivitiesFromSlackBatch(ctx, *msg); err != nil {
		log.Fatalf("handler.FetchActivitiesFromSlackBatch failed.(err=%+v)", err)
	}
}
