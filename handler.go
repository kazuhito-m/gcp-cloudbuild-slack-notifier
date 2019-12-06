package main

import (
	"context"
	"encoding/json"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/pubsub"
	"log"
)

func HelloFromPubSub(ctx context.Context, message pubsub.PubSubMessage) error {
	var decode_data cloudbuild.CloudBuildResult
	if err := json.Unmarshal(message.Data, &decode_data); err != nil {
		log.Fatal(err)
		return nil
	}

	text := string(message.Data)
	log.Println("取得できたData部:" + text)

	return nil
}
