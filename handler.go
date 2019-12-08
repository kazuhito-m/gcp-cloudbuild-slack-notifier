package handler

import (
	"context"
	"encoding/json"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/notify"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/pubsub"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
	"log"
)

func PubSubHandlerForCloudBuild(ctx context.Context, message pubsub.PubSubMessage) error {
	var decode_data cloudbuild.CloudBuildResult
	if err := json.Unmarshal(message.Data, &decode_data); err != nil {
		log.Fatal(err)
		return nil
	}

	text := string(message.Data)
	log.Println("取得できたData部:" + text)

	config, ok := config.Load()
	if !ok {
		log.Fatalln("環境変数から設定が取得できませんでした。終了します。")
		return nil
	}
	log.Println(config.SlackURL)

	notify := notify.CreateNotify()

	result := slack.SendNotify(notify, config)

	log.Println(endInfo(result))

	return nil
}

func endInfo(result bool) string {
	resText := "失敗"
	if result {
		resText = "成功"
	}
	return "Slack通知が" + resText + "しました。"
}
