package firestore_repository

import (
	"context"
	"log"

	"example.com/asakatsu-app/infrastructure/client"

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
	firestoreHandler *client.FirestoreHandler,
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
		log.Printf("Activity document is not found(err=%+v", err)
		return nil, err
	} else if err != nil {
		log.Printf("Activity document get failed(err=%+v)", err)
		return nil, err
	}

	activityDoc := &firestore_entity.ActivityDoc{
		ID: docID,
	}

	if err = dataSnap.DataTo(&activityDoc.Field); err != nil {
		log.Printf("Can not convert to activity doc entity.(err=%+v)", err)
		return nil, err
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
		log.Printf("Actibities get failed(err=%+v)", err)
		return nil, err
	}
	if len(dataSpans) == 0 {
		log.Print("ActivityRepository.GetAllBySlackUID result count is zero")
		return nil, nil
	}

	var docFieldList []firestore_entity.ActivityField
	for _, dataSnap := range dataSpans {
		docFiled := new(firestore_entity.ActivityField)
		if err = dataSnap.DataTo(docFiled); err != nil {
			log.Printf("Can not convert to activity field entity(err=%+v)", err)
			return nil, err
		}

		docFieldList = append(docFieldList, *docFiled)
	}

	return docFieldList, nil
}

// Set は、activitiesテーブルにドキュメントを保存します。
// see https://pkg.go.dev/cloud.google.com/go/firestore#DocumentRef.Set
func (r *ActivityRepository) Set(activityDoc firestore_entity.ActivityDoc) error {
	if _, err := r.collection.Doc(activityDoc.ID).Set(r.ctx, activityDoc.Field); err != nil {
		log.Printf("Can not set activity doc to firestore.(err=%+v)", err)
		return err
	}

	return nil
}
