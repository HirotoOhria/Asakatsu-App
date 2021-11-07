package injector

import (
	"context"

	"example.com/asakatsu-app/handler/function_handler"
)

func InjectFetchActivitiesFromSlackHandler(
	ctx context.Context,
) *function_handler.FetchActivitiesFromSlackHandler {
	return function_handler.NewFetchActivitiesFromSlackHandler(InjectFetchActivitiesFromSlackBatchUsecase(ctx))
}

func InjectGetActivitiesFromSlackUidHandler(
	ctx context.Context,
) *function_handler.GetActivitiesFromSlackUidHandler {
	return function_handler.NewGetActivitiesFromSlackUidHandler(InjectGetActivitiesFromSlackUidUsecase(ctx))
}
