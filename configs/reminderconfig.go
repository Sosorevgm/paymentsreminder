package configs

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

const (
	webHookEnv      = "WEB_HOOK_URL"
	messageTitle    = "MESSAGE_TITLE"
	messageSubtitle = "MESSAGE_SUBTITLE"
)

type ReminderConfig struct {
	WebHook         string
	MessageTitle    string
	MessageSubtitle string
}

func GetConfig() (ReminderConfig, error) {
	err := godotenv.Load("paymentsreminder.env")
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

	return ReminderConfig{webHook, title, subtitle}, err
}
