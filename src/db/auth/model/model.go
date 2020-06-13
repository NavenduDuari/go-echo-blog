package model

type OtpAuth struct {
	Id        string `json:"id"`
	UserId    string `json:"userid"`
	OTP       string `json:"otp"`
	ExpiresAt string `json:"expires_at"`
	// CreatedAt string `json:"created-at"`
}

type QrAuth struct {
	Id        string `json:"id"`
	UserId    string `json:"userid"`
	QrCode    string `json:"qr_code"`
	ExpiresAt string `json:"expires_at"`
	// CreatedAt string `json:"created-at"`
}

type SubscribeSnsTopicInput struct {
	Endpoint string `json:"endpoint"`
	Protocol string `json:"protocol"`
	TopicArn string `json:"topic-arn"`
}

type PublishSnsTopicInput struct {
	Message  OtpAuth `json:"message"`
	TopicArn string  `json:"topic-arn"`
}
