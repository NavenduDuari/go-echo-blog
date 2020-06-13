package auth

import (
	"fmt"
	"log"

	"github.com/NavenduDuari/go-echo-blog/src/common/utils"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/internal/otpmanager"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/internal/store"
	"github.com/NavenduDuari/go-echo-blog/src/db/auth/model"
)

func CreateTopic(topicName string) (string, error) {
	sess := utils.GetSnsSession()
	topicArn, err := otpmanager.CreateSnsTopic(sess, topicName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return topicArn, nil
}

func DeleteTopic(topicArn string) error {
	sess := utils.GetSnsSession()
	err := otpmanager.DeleteSnsTopic(sess, topicArn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SubscribeTopic(subscribeInput model.SubscribeSnsTopicInput) error {
	sess := utils.GetSnsSession()
	_, err := otpmanager.SubscribeSnsTopic(sess, subscribeInput)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Publish(publishInput model.PublishSnsTopicInput) error {
	sess := utils.GetSnsSession()
	_, err := otpmanager.PublishToSnsTopic(sess, publishInput)
	if err != nil {
		fmt.Println(err)
		return err
	}

	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	if err = store.InsertOtp(db, publishInput.Message); err != nil {
		return err
	}
	return nil
}

func FetchOtp(userId string) (model.OtpAuth, error) {
	db, err := store.GetPostgressDB()
	if err != nil {
		log.Fatal("GetPostgressDB failed", err)
	}
	otpAuth, err := store.FetchOtpByUserId(db, userId)
	if err != nil {
		return otpAuth, err
	}
	return otpAuth, nil
}
