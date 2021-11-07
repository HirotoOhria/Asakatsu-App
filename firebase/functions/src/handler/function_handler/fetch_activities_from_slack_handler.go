package function_handler

import (
	"log"

	"example.com/asakatsu-app/usecase"
)

type FetchActivitiesFromSlackHandler struct {
	*usecase.FetchActivitiesFromSlackUsecase
}

func NewFetchActivitiesFromSlackHandler(
	usecase *usecase.FetchActivitiesFromSlackUsecase,
) *FetchActivitiesFromSlackHandler {
	return &FetchActivitiesFromSlackHandler{
		FetchActivitiesFromSlackUsecase: usecase,
	}
}

// FetchActivitiesFromSlackBatch は、FetchActivitiesFromSlackBatch を実行します。
func (h *FetchActivitiesFromSlackHandler) Exec() error {
	log.Print("run: handler.FetchActivitiesFromSlackBatch()")

	if err := h.FetchActivitiesFromSlackUsecase.Exec(); err != nil {
		log.Printf("FetchActivitiesFromSlackUsecase.Exec failed(err=%+v)", err)
		return err
	}

	return nil
}
