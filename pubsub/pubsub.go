package pubsub

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"regexp"
	"strings"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Info struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

func (i PubSubMessage) IsCloudBuildMessage() bool {
	messageJson := string(i.Data)
	statusText := PickUpStatusText(messageJson)
	return cloudbuild.InCloudBuildStatus(statusText)
}

func PickUpStatusText(json string) string {
	r := regexp.MustCompile(`("status"\ *:\ *"[A-Z]+")`)
	hits := r.FindAllStringSubmatch(json, -1)
	text := hits[0][0]
	text = erase(text, `"status"`)
	text = erase(text, ":")
	text = erase(text, `"`)
	return strings.TrimSpace(text)
}

func erase(target string, word string) string {
	return strings.Replace(target, word, "", -1)
}

func (i PubSubMessage) DataText() string {
	return string(i.Data)
}
