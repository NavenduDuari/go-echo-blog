package otpmanager

import (
	"encoding/json"
	"fmt"

	"github.com/NavenduDuari/go-echo-blog/src/db/auth/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func CreateSnsTopic(sess *session.Session, topicName string) (string, error) {
	svc := sns.New(sess)
	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(topicName),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(*result.TopicArn)
	return *result.TopicArn, nil
}

func DeleteSnsTopic(sess *session.Session, topicArn string) error {
	svc := sns.New(sess)
	_, err := svc.DeleteTopic(&sns.DeleteTopicInput{
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SubscribeSnsTopic(sess *session.Session, subscribeInput model.SubscribeSnsTopicInput) (string, error) {
	svc := sns.New(sess)
	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              aws.String(subscribeInput.Endpoint),
		Protocol:              aws.String(subscribeInput.Protocol),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              aws.String(subscribeInput.TopicArn),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(*result.SubscriptionArn)
	return *result.SubscriptionArn, nil
}

func PublishToSnsTopic(sess *session.Session, publishInput model.PublishSnsTopicInput) (string, error) {
	svc := sns.New(sess)
	message, _ := json.Marshal(publishInput.Message)
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(message)),
		TopicArn: aws.String(publishInput.TopicArn),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return *result.MessageId, nil
}
