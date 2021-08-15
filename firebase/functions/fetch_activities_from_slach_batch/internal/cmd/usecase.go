package cmd

import (
	"log"
	"time"

	"example.com/fetch-activities-from-slack-batch/internal/firebase/firestore/firestore_entity"
	"example.com/fetch-activities-from-slack-batch/internal/firebase/firestore/firestore_repository"
	"example.com/fetch-activities-from-slack-batch/internal/slack/slack_repository"
	"example.com/fetch-activities-from-slack-batch/internal/utils"

	"github.com/slack-go/slack"
)

type Usecase struct {
	asakatsuRepository *slack_repository.AsakatsuRepository
	activityRepository *firestore_repository.ActivityRepository
}

func NewUsecase(
	asakatsuRepository *slack_repository.AsakatsuRepository,
	activityRepository *firestore_repository.ActivityRepository,
) *Usecase {
	return &Usecase{
		asakatsuRepository,
		activityRepository,
	}
}

func (u *Usecase) Exec() {
	log.Print("run: cmd.Exec()")

	startMsgs := u.getStartMsgs()

	for _, startMsg := range startMsgs {
		u.saveActivity(startMsg)
	}
}

func (u *Usecase) getStartMsgs() []slack.Message {
	response, err := u.asakatsuRepository.GetYesterdayConversationHistory()
	if err != nil {
		log.Fatalf("GetYesterdayConversationHistory failed.(err=%+v)", err)
	}

	return utils.FindStartConversations(response.Messages)
}

func (u *Usecase) saveActivity(startMsg slack.Message) {
	startTime, err := u.buildStartTime(startMsg)
	if err != nil {
		log.Fatalf("build startTime failed.(err=%+v)", err)
	}

	endTime, err := u.buildEndTime(startMsg)
	if err != nil {
		log.Fatalf("build endTime failed.(err=%+v)", err)
	}

	activityData := &firestore_entity.ActivityData{
		SlackUID:  startMsg.User,
		StartTime: startTime,
		EndTime:   endTime,
	}
	activityDoc := firestore_entity.NewActivityDoc(startTime, activityData)

	err = u.activityRepository.Set(*activityDoc)
	if err != nil {
		log.Fatalf("activityRepository set failed.(err=%+v)", err)
	}
}

func (u *Usecase) buildStartTime(startMsg slack.Message) (t time.Time, err error) {
	t, err = utils.ParseTime(startMsg.Timestamp)
	if err != nil {
		log.Fatalf("ParseTime failed.(err=%+v)", err)
		return t, err
	}

	return t, nil
}

func (u *Usecase) buildEndTime(startMsg slack.Message) (t time.Time, err error) {
	replieMsgs, err := u.asakatsuRepository.GetConversationReplies(startMsg.Timestamp)
	if err != nil {
		log.Fatalf("GetConversationReplies failed.(err=%+v)", err)
		return t, err
	}

	endMsgs := utils.FindEndConversations(replieMsgs)

	if len(endMsgs) == 0 {
		return t, nil
	}

	t, err = utils.ParseTime(endMsgs[0].Timestamp)
	if err != nil {
		log.Fatalf("ParseTime failed.(err=%+v)", err)
		return t, err
	}

	return t, nil
}
