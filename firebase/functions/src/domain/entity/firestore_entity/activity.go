package firestore_entity

import "time"

const idTimeFormat = time.RFC3339

type ActivityDoc struct {
	ID    string
	Field ActivityField
}

type ActivityField struct {
	SlackUID  string
	StartTime *time.Time
	EndTime   *time.Time
}

func NewActivityDoc(t time.Time, field ActivityField) *ActivityDoc {
	id := t.Format(idTimeFormat)

	return &ActivityDoc{
		ID:    id,
		Field: field,
	}
}
