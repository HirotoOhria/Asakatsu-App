package injector

import (
	"context"

	"example.com/asakatsu-app/usecase/batch_usecase"
)

func InjectFetchActivitiesFromSlackBatchUsecase(ctx context.Context) *batch_usecase.FetchActivitiesFromSlackUsecase {
	return batch_usecase.NewFetchActivitiesFromSlackUsecase(InjectAsakatsuRepository(), InjectActivityRepostiroy(ctx))
}
