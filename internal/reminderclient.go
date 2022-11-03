package internal

import (
	"bytes"
	"net/http"
)

type ReminderClient struct {
	client http.Client
}

func NewReminderClient() ReminderClient {
	return ReminderClient{http.Client{}}
}

func (reminderClient ReminderClient) SendMessage(
	webhookUrl string,
	title string,
	subtitle string,
) error {
	msg, err := NewRemindMessage(title, subtitle)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}
	_, err = reminderClient.client.Do(request)
	return err
}
