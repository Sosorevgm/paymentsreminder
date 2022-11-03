package main

import (
	"log"
	"paymentsreminder/configs"
	"paymentsreminder/internal"
	"time"
)

func main() {
	client := internal.NewReminderClient()
	config, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	for {
		weekDay := time.Now().Weekday()
		hours, minutes, _ := time.Now().Clock()

		isNeedNotify := internal.CheckTime(
			weekDay,
			hours,
			minutes,
			config.NotificationDays,
			config.NotificationHour,
			config.NotificationMinute,
		)

		if isNeedNotify {
			err = client.SendMessage(
				config.WebHook,
				config.MessageTitle,
				config.MessageSubtitle,
			)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Hour)
		} else {
			time.Sleep(time.Second)
		}
	}
}
