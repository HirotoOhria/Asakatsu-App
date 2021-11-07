package usecase

import (
	"log"
	"time"

	"example.com/asakatsu-app/domain/domain_object/slack_domain_object"
	"example.com/asakatsu-app/domain/entity/firestore_entity"
	"example.com/asakatsu-app/domain/repository/firestore_repository"
	"example.com/asakatsu-app/domain/repository/slack_repository"
)

// FetchActivitiesFromSlackUsecase is usecase
type FetchActivitiesFromSlackUsecase struct {
	AsakatsuRepository *slack_repository.AsakatsuRepository
	ActivityRepository *firestore_repository.ActivityRepository
}

// NewFetchActivitiesFromSlackUsecase is connstractor
func NewFetchActivitiesFromSlackUsecase(
	asakatsuRepository *slack_repository.AsakatsuRepository,
	activityRepository *firestore_repository.ActivityRepository,
) *FetchActivitiesFromSlackUsecase {
	return &FetchActivitiesFromSlackUsecase{
		asakatsuRepository,
		activityRepository,
	}
}

// Exec は、FetchActivitiesFromSlackBatchの処理を実行します。
// 昨日のSlackメッセージの中から :start: が含まれるメッセージを取得し、それぞれのアクティビティを保存します。
func (u *FetchActivitiesFromSlackUsecase) Exec() error {
	log.Print("run: FetchActivitiesFromSlackUsecase.Exec()")

	yesterdaySlackMsgs, err := u.AsakatsuRepository.GetYesterdayConversationHistory()
	if err != nil {
		log.Printf("AsakatsuRepository.GetYesterdayConversationHistory failed.(err=%+v)", err)
		return err
	}

	startSlackMsgs := slack_domain_object.NewSlackConversation(yesterdaySlackMsgs).FindStartSlackMsgs()
	if startSlackMsgs == nil {
		log.Print("yesterday's start slack message count is zero")
		return nil
	} else {
		log.Printf("yesterday's start slack message count is %+v", len(startSlackMsgs))
	}

	for _, startSlackMsg := range startSlackMsgs {
		if err := u.saveActivityFromStartSlackMsg(startSlackMsg); err != nil {
			return err
		}
	}

	return nil
}

// saveActivityFromStartSlackMsg は、Slackのスタートメッセージからアクティビティを保存します。
func (u *FetchActivitiesFromSlackUsecase) saveActivityFromStartSlackMsg(
	startSlackMsg slack_domain_object.SlackStartMsg,
) error {
	replyMsgs, err := u.AsakatsuRepository.GetConversationReplies(startSlackMsg.Content.Timestamp)
	if err != nil {
		log.Printf("AsakatsuRepository.GetConversationReplies failed.(err=%+v)", err)
		return err
	}

	var endTime *time.Time
	if slackEndMsg := slack_domain_object.NewSlackReplies(replyMsgs).FindEndMsg(); slackEndMsg == nil {
		endTime = nil
	} else {
		endTime = slackEndMsg.GetTime()
	}

	activityField := &firestore_entity.ActivityField{
		SlackUID:  startSlackMsg.Content.User,
		StartTime: startSlackMsg.GetTime(),
		EndTime:   endTime,
	}
	activityDoc := firestore_entity.NewActivityDoc(*startSlackMsg.GetTime(), *activityField)

	if err = u.ActivityRepository.Set(*activityDoc); err != nil {
		log.Printf("ActivityRepository.Set failed(err=%+v)", err)
		return err
	}

	return nil
}
