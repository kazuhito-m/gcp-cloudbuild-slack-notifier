package config

import (
	"os"
)

type Config struct {
	SlackURL     string
	SlackChannel string
}

func Load() (Config, bool) {
	config := Config{}
	config.SlackURL = os.Getenv("GCSN_SLACK_URL")
	config.SlackChannel = os.Getenv("GCSN_SLACK_CHANNEL")
	if config.SlackURL == "" {
		config.SlackChannel = "#general"
	}
	return config, config.SlackURL != ""
}
