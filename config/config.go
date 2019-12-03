package config

import (
	"os"
)

type Config struct {
	SlackURL string
}

func Load() (Config, bool) {
	config := Config{}
	config.SlackURL = os.Getenv("GCSN_SLACK_URL")
	return config, config.SlackURL != ""
}
