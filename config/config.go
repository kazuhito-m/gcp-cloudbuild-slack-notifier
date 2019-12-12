package config

import (
	"os"
	"strconv"
)

type Config struct {
	SlackURL           string
	SlackChannel       string
	StartNotifyDisable bool
	EndNotifyDisable   bool
	TriggerSettings    []TriggerSetting
}

type TriggerSetting struct {
	TriggerID string
	Disable   bool
	AliasName string
}

func Load() (Config, bool) {
	config := Config{}

	config.SlackURL = os.Getenv("GCSN_SLACK_URL")

	config.SlackChannel = os.Getenv("GCSN_SLACK_CHANNEL")
	if config.SlackURL == "" {
		config.SlackChannel = "#general"
	}

	config.StartNotifyDisable = false
	b, err := strconv.ParseBool(os.Getenv("GCSN_START_NOTIFY_DISABLE"))
	if err == nil {
		config.StartNotifyDisable = b
	}

	config.EndNotifyDisable = false
	b, err = strconv.ParseBool(os.Getenv("GCSN_END_NOTIFY_DISABLE"))
	if err == nil {
		config.EndNotifyDisable = b
	}

	config.TriggerSettings = createTriggerSettings()

	return config, config.SlackURL != ""
}

func (i Config) TriggerSettingOf(triggerID string) (TriggerSetting, bool) {
	for _, i := range i.TriggerSettings {
		if i.TriggerID == triggerID {
			return i, true
		}
	}
	return TriggerSetting{}, false
}
