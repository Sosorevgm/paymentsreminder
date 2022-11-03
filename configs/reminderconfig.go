package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	webHookEnv         = "WEB_HOOK_URL"
	messageTitle       = "MESSAGE_TITLE"
	messageSubtitle    = "MESSAGE_SUBTITLE"
	notificationHour   = "NOTIFICATION_HOUR"
	notificationMinute = "NOTIFICATION_MINUTE"
	notificationDays   = "NOTIFICATION_DAYS"
)

type ReminderConfig struct {
	WebHook            string
	MessageTitle       string
	MessageSubtitle    string
	NotificationHour   int
	NotificationMinute int
	NotificationDays   []time.Weekday
}

func GetConfig() (ReminderConfig, error) {
	err := godotenv.Load("configexample.env")
	if err != nil {
		return ReminderConfig{}, err
	}
	webHook, exist := os.LookupEnv(webHookEnv)
	if exist == false {
		return ReminderConfig{}, errors.New("empty web hook environment")
	}
	title, exist := os.LookupEnv(messageTitle)
	if exist == false {
		return ReminderConfig{}, errors.New("empty message title environment")
	}
	subtitle, exist := os.LookupEnv(messageSubtitle)
	if exist == false {
		return ReminderConfig{}, errors.New("empty message subtitle environment")
	}
	hourStr, exist := os.LookupEnv(notificationHour)
	if exist == false {
		return ReminderConfig{}, errors.New("empty notification hour environment")
	}
	minuteStr, exist := os.LookupEnv(notificationMinute)
	if exist == false {
		return ReminderConfig{}, errors.New("empty notification minute environment")
	}
	daysStr, exist := os.LookupEnv(notificationDays)
	if exist == false {
		return ReminderConfig{}, errors.New("empty notification days environment")
	}
	hour, err := strconv.Atoi(hourStr)
	if err != nil || hour < 0 || hour > 23 {
		return ReminderConfig{}, errors.New("wrong notification hour format")
	}
	minute, err := strconv.Atoi(minuteStr)
	if err != nil || minute < 0 || minute > 59 {
		return ReminderConfig{}, errors.New("wrong notification minute format")
	}
	daysStrArr := strings.Split(daysStr, ",")
	days := make([]time.Weekday, len(daysStrArr))
	for index, dayStr := range daysStrArr {
		dayInt, err := strconv.Atoi(dayStr)
		if err != nil || dayInt < 0 || dayInt > 6 {
			return ReminderConfig{}, errors.New("wrong notification days format")
		}
		days[index] = time.Weekday(dayInt)
	}
	return ReminderConfig{
		webHook,
		title,
		subtitle,
		hour,
		minute,
		days,
	}, err
}
