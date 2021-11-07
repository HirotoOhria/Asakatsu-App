package injector

import (
	"context"

	"example.com/asakatsu-app/usecase"
)

func InjectFetchActivitiesFromSlackBatchUsecase(
	ctx context.Context,
) *usecase.FetchActivitiesFromSlackUsecase {
	return usecase.NewFetchActivitiesFromSlackUsecase(
		InjectAsakatsuRepository(),
		InjectActivityRepostiroy(ctx),
	)
}

func InjectGetActivitiesFromSlackUidUsecase(
	ctx context.Context,
) *usecase.GetActivitiesFromSlackUidUsecase {
	return usecase.NewGetActivitiesFromSlackUidUsecase(InjectActivityRepostiroy(ctx))
}
