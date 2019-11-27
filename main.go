package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
)

func main() {
	fmt.Printf("%v\n", uuid.Must(uuid.NewV4()))
	slack.Sample()
}
