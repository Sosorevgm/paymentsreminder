package main

import (
	"bytes"
	"log"
	"net/http"
	"paymentsreminder/configs"
	"paymentsreminder/internal"
	"time"
)

func main() {
	client := &http.Client{}
	config, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	for {
		weekDay := time.Now().Weekday()
		hours, minutes, _ := time.Now().Clock()

		isNeedNotify := checkTime(weekDay, hours, minutes)

		if isNeedNotify {
			err = sendRequest(*client, config)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Hour)
		} else {
			time.Sleep(time.Second)
		}
	}
}

func checkTime(weekDay time.Weekday, hours int, minutes int) bool {
	if weekDay == time.Monday || weekDay == time.Tuesday || weekDay == time.Wednesday || weekDay == time.Thursday {
		if hours == 11 {
			if minutes == 30 {
				return true
			}
		}
	}
	return false
}

func sendRequest(client http.Client, config configs.ReminderConfig) error {
	msg, err := internal.NewRemindMessage(config.MessageTitle, config.MessageSubtitle)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", config.WebHook, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}

	_, err = client.Do(request)
	return err
}
