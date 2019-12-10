package handler

import (
	"context"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/notify"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/pubsub"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
	"log"
)

func PubSubHandlerForCloudBuild(ctx context.Context, message pubsub.PubSubMessage) error {
	text := string(message.Data)

	if !message.IsCloudBuildMessage() {
		log.Println("CloudBuild以外の通知だったのでスキップしました。:" + text)
		return nil
	}

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
