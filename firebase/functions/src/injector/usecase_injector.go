package injector

import (
	"context"

	"example.com/asakatsu-app/usecase/batch_usecase"
	"example.com/asakatsu-app/usecase/function_usecase"
)

func InjectFetchActivitiesFromSlackBatchUsecase(
	ctx context.Context,
) *batch_usecase.FetchActivitiesFromSlackUsecase {
	return batch_usecase.NewFetchActivitiesFromSlackUsecase(
		InjectAsakatsuRepository(),
		InjectActivityRepostiroy(ctx),
	)
}

func InjectGetActivitiesFromSlackUidUsecase(
	ctx context.Context,
) *function_usecase.GetActivitiesFromSlackUidUsecase {
	return function_usecase.NewGetActivitiesFromSlackUidUsecase(InjectActivityRepostiroy(ctx))
}
