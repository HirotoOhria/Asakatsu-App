package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/joho/godotenv"

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

	w := httptest.NewRecorder()
	r := &http.Request{
		URL: makeUrlForLocal(),
	}

	output := handler.GetActivitiesFromSlackUidFunction(w, r)
	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatalf("json.Marshal failed(err=%+v)", err)
	}

	fmt.Print(string(outputJson))
}

// makeUrlForLocal は、ローカル実行用のURLを生成します
func makeUrlForLocal() *url.URL {
	urlForLocal, err := url.Parse("https://example.com/?slack_uid=" + getSlackUidByArg())
	if err != nil {
		log.Fatalf("url.Parse failed(err=%+v)", err)
	}

	return urlForLocal
}

// getSlackUidByArg は、go run コマンドの第一引数からSlakcUidを取得します
func getSlackUidByArg() string {
	flag.Parse()
	return flag.Arg(0)
}
