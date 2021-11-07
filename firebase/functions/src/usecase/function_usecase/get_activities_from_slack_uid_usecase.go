package function_usecase

import (
	"log"

	"example.com/asakatsu-app/domain/entity/firestore_entity"
	"example.com/asakatsu-app/domain/repository/firestore_repository"
)

// GetActivitiesFromSlackUidUsecase は、GetActivitiesFromSlackUidFunction のユースケースです。
type GetActivitiesFromSlackUidUsecase struct {
	*firestore_repository.ActivityRepository
}

// NewGetActivitiesFromSlackUidUsecase は、コンストラクタです。
func NewGetActivitiesFromSlackUidUsecase(
	activityRepository *firestore_repository.ActivityRepository,
) *GetActivitiesFromSlackUidUsecase {
	return &GetActivitiesFromSlackUidUsecase{
		ActivityRepository: activityRepository,
	}
}

func (u *GetActivitiesFromSlackUidUsecase) Exec(slackUID string) ([]firestore_entity.ActivityField, error) {
	log.Print("run: GetActivitiesFromSlackUidUsecase.Exec()")

	activityFieldList, err := u.ActivityRepository.GetAllBySlackUID(slackUID)
	if err != nil {
		log.Printf("ActivityRepository.GetAllBySlackUID failed.(err=%+v)", err)
		return nil, err
	}
	if activityFieldList == nil {
		log.Print("ActivityRepository.GetAllBySlackUID result count is zero")
	} else {
		log.Printf("ActivityRepository.GetAllBySlackUID result count is %+v", len(activityFieldList))
	}

	return activityFieldList, err
}
