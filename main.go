package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/notify"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
	"log"
)

func main() {
	fmt.Printf("%v\n", uuid.Must(uuid.NewV4()))
	config, ok := config.Load()
	if !ok {
		log.Fatalln("環境変数から設定が取得できませんでした。終了します。")
		return
	}
	log.Println(config.SlackURL)

	notify := notify.CreateNotify()

	result := slack.SendNotify(notify, config)

	log.Println(endInfo(result))
}

func endInfo(result bool) string {
	resText := "失敗"
	if result {
		resText = "成功"
	}
	return "Slack通知が" + resText + "しました。"
}
