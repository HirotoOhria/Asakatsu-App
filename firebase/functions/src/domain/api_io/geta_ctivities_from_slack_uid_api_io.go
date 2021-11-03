package api_io

import (
	"time"

	"example.com/asakatsu-app/domain/entity/firestore_entity"
)

// GetActivitiesFromSlackUidInput は、GetActivitiesFromSlackUidFunctino のインプット
type GetActivitiesFromSlackUidInput struct {
	SlackUID string
}

// GetActivitiesFromSlackUidOutput は、GetActivitiesFromSlackUidFunctino のアウトプット
type GetActivitiesFromSlackUidOutput struct {
	ActivityFieldOutputs []ActivityFieldOutput `json:"activityFields"`
}

// ActivityFieldOutput は、ActivityField エンティティのアウトプット
type ActivityFieldOutput struct {
	SlackUID  string     `json:"slackUid"`
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
}

// NewGetActivitiesFromSlackUidOutput は、コンストラクタ
func NewGetActivitiesFromSlackUidOutput(
	activityFields []firestore_entity.ActivityField,
) *GetActivitiesFromSlackUidOutput {
	var activityFieldOutputs []ActivityFieldOutput
	for _, activityField := range activityFields {
		activityFieldOutputs = append(activityFieldOutputs, ActivityFieldOutput{
			SlackUID:  activityField.SlackUID,
			StartTime: activityField.StartTime,
			EndTime:   activityField.EndTime,
		})
	}

	return &GetActivitiesFromSlackUidOutput{
		ActivityFieldOutputs: activityFieldOutputs,
	}
}
