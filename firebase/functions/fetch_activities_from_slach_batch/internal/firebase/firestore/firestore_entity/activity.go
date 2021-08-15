package firestore_entity

import "time"

const format = time.RFC3339

type ActivityDoc struct {
	ID   string
	Data ActivityData
}

type ActivityData struct {
	SlackUID  string
	StartTime time.Time
	EndTime   time.Time
	// ActivityWorks []string
}

func NewActivityDoc(t time.Time, data *ActivityData) *ActivityDoc {
	id := t.Format(format)

	return &ActivityDoc{
		ID:   id,
		Data: *data,
	}
}
