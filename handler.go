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
	if !message.IsCloudBuildMessage() {
		log.Println("CloudBuild以外の通知だったのでスキップしました。:" + message.DataText())
		return nil
	}
	result := message.ToCloudBuildResult()
	if !result.IsStart() && !result.IsEnd() {
		log.Println("CloudBuildの開始・終了通知ではなかったためスキップしました。:" + message.DataText())
		return nil
	}

	config, ok := config.Load()
	if !ok {
		log.Fatalln("環境変数から設定が取得できませんでした。終了します。")
		return nil
	}
	log.Println(config.SlackURL)

	notify := notify.CreateNotify(result)

	okOrNg := slack.SendNotify(notify, config)

	log.Println(endInfo(okOrNg))

	return nil
}

func endInfo(result bool) string {
	resText := "失敗"
	if result {
		resText = "成功"
	}
	return "Slack通知が" + resText + "しました。"
}
