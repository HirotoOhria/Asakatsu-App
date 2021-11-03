package firestore_repository

import (
	"context"
	"fmt"

	"example.com/asakatsu-app/infrastructure"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"example.com/asakatsu-app/domain/entity/firestore_entity"
)

// activitiesCollection は、Firesotoreのアクティビティテーブルの名前です。
const activitiesCollection = "activities"

// ActivityRepository は、アクティビティテーブルのリポジトリです。
type ActivityRepository struct {
	ctx        context.Context
	collection *firestore.CollectionRef
}

// ActivityRepository は、コンストラクタです。
func NewActivityRepostitory(
	ctx context.Context,
	firestoreHandler *infrastructure.FirestoreHandler,
) *ActivityRepository {
	return &ActivityRepository{
		ctx:        ctx,
		collection: firestoreHandler.DB.Collection(activitiesCollection),
	}
}

// Get は、actibitiesテーブルからドキュメントを取得します。
// 対象のドキュメントが存在しない場合、nilを返します。
// see https://pkg.go.dev/cloud.google.com/go/firestore#DocumentRef.Get
func (r *ActivityRepository) Get(docID string) (*firestore_entity.ActivityDoc, error) {
	dataSnap, err := r.collection.Doc(docID).Get(r.ctx)
	if status.Code(err) == codes.NotFound {
		return nil, fmt.Errorf("docment is not found.(err=%+v)", err)
	} else if err != nil {
		return nil, fmt.Errorf("get dockment failed(err=%+v)", err)
	}

	activityDoc := &firestore_entity.ActivityDoc{
		ID: docID,
	}

	if err = dataSnap.DataTo(&activityDoc.Field); err != nil {
		return nil, fmt.Errorf("can not convert to activity doc entity.(err=%+v)", err)
	}

	return activityDoc, nil
}

// GetAllBySlackUID は、SlackUIDが一致するすべてのアクティビティのフィールドを取得します。
// 対象のドキュメントが存在しない場合、nilを返します。
// see https://pkg.go.dev/cloud.google.com/go/firestore#Query.Where
func (r *ActivityRepository) GetAllBySlackUID(slackUID string) ([]firestore_entity.ActivityField, error) {
	dataSpans, err := r.collection.
		Where("SlackUID", "==", slackUID).
		OrderBy("StartTime", firestore.Desc).
		Documents(r.ctx).
		GetAll()
	if err != nil {
		return nil, fmt.Errorf("can not get documents by slackUID.(err=%+v)", err)
	}

	var docFieldList []firestore_entity.ActivityField
	for _, dataSnap := range dataSpans {
		docFiled := new(firestore_entity.ActivityField)
		if err = dataSnap.DataTo(docFiled); err != nil {
			return nil, fmt.Errorf("can not convert to activity field entity.(err=%+v)", err)
		}

		docFieldList = append(docFieldList, *docFiled)
	}

	return docFieldList, nil
}

// Set は、activitiesテーブルにドキュメントを保存します。
// see https://pkg.go.dev/cloud.google.com/go/firestore#DocumentRef.Set
func (r *ActivityRepository) Set(activityDoc firestore_entity.ActivityDoc) error {
	if _, err := r.collection.Doc(activityDoc.ID).Set(r.ctx, activityDoc.Field); err != nil {
		return fmt.Errorf("can not set activity doc to firestore.(err=%+v)", err)
	}

	return nil
}
