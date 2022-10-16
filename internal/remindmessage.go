package internal

import "encoding/json"

type RemindMessage struct {
	Type       string    `json:"@type"`
	Context    string    `json:"@context"`
	ThemeColor string    `json:"themeColor"`
	Summary    string    `json:"summary"`
	Sections   []Section `json:"sections"`
}

type Section struct {
	ActivityTitle    string `json:"activityTitle"`
	ActivitySubtitle string `json:"activitySubtitle"`
	ActivityImage    string `json:"activityImage"`
	Facts            []Fact `json:"facts"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewRemindMessage(title string, subtitle string) ([]byte, error) {
	msg := RemindMessage{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: "#00FF00",
		Summary:    "Payments reminder",
		Sections: []Section{
			{
				ActivityTitle:    title,
				ActivitySubtitle: subtitle,
			},
		},
	}
	res, err := json.Marshal(msg)
	return res, err
}
