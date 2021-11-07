package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"example.com/asakatsu-app/domain/api_io"
	"example.com/asakatsu-app/handler"
)

func init() {
	log.Print("run: main.init()")

	initDotenv()
}

// initDotenv は、.env ファイルの環境変数を読み込みます。
func initDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error: loading .env file")
	}
}

// ex. go run main.go ${slack_uid}
func main() {
	log.Print("run: main.main()")

	flag.Parse()
	input := &api_io.GetActivitiesFromSlackUidInput{
		SlackUID: flag.Arg(0),
	}

	output, err := handler.GetActivitiesFromSlackUidFunction(input)
	if err != nil {
		log.Fatalf("handler.GetActivitiesFromSlackUidFunction failed(err=%+v)", err)
	}
	if output == nil {
		log.Fatal("handler.GetActivitiesFromSlackUidFunction not found")
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatalf("json.Marshal failed(err=%+v)", err)
	}

	fmt.Printf(string(outputJson))
}
