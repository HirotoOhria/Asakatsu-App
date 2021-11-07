package function_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"example.com/asakatsu-app/domain/api_io"
	"example.com/asakatsu-app/usecase"
)

type GetActivitiesFromSlackUidHandler struct {
	*usecase.GetActivitiesFromSlackUidUsecase
}

func NewGetActivitiesFromSlackUidHandler(
	usecase *usecase.GetActivitiesFromSlackUidUsecase,
) *GetActivitiesFromSlackUidHandler {
	return &GetActivitiesFromSlackUidHandler{
		GetActivitiesFromSlackUidUsecase: usecase,
	}
}

// GetActivitiesFromSlackUidHandler は、GetActivitiesFromSlackUidUsecase を実行します。
// endpoint https://asia-northeast1-asakatsu-app-d6f28.cloudfunctions.net/GetActivitiesFromSlackUidFunction?slack_uid=${slack_uid}
func (h *GetActivitiesFromSlackUidHandler) Exec(
	resWiter http.ResponseWriter,
	req *http.Request,
) *api_io.GetActivitiesFromSlackUidOutput {
	// TODO 404 レスポンスを返す
	log.Print("run: handler.GetActivitiesFromSlackUidFunction()")

	input, err := getInputByReq(req)
	if err != nil {
		log.Printf("GetActivitiesFromSlackUidHandler.getInutByReq failed(err=%+v)", err)
		return nil
	}

	output, err := h.GetActivitiesFromSlackUidUsecase.Exec(input.SlackUID)
	if err != nil {
		log.Printf("GetActivitiesFromSlackUidUsecase.Exec failed(err=%+v)", err)
		return nil
	}

	err = writeResFromOutput(resWiter, output)
	if err != nil {
		log.Printf("GetActivitiesFromSlackUidHandler.writeResFromOutput failed(err=%+v)", err)
		return nil
	}

	return output
}

func getInputByReq(r *http.Request) (*api_io.GetActivitiesFromSlackUidInput, error) {
	slackUid := r.URL.Query().Get("slack_uid")
	if slackUid == "" {
		log.Print("slack_uid query param not found")
		return nil, errors.New("slack_uid query param not found")
	}

	input := &api_io.GetActivitiesFromSlackUidInput{
		SlackUID: slackUid,
	}
	log.Printf("input: %+v", input)

	return input, nil
}

func writeResFromOutput(w http.ResponseWriter, output *api_io.GetActivitiesFromSlackUidOutput) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Printf("json.Marshal failed(err%+v)", err)
		return fmt.Errorf("json.Marshal failed(err%+v)", err)
	}

	_, err = fmt.Fprintf(w, string(outputJson))
	if err != nil {
		log.Printf("fmt.Fprintf failed(err=%+v)", err)
		return fmt.Errorf("fmt.Fprintf failed(err=%+v)", err)
	}

	return nil
}
