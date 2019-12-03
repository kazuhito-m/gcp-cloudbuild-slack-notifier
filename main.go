package main

import (
	"fmt"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"log"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Printf("%v\n", uuid.Must(uuid.NewV4()))
	config, ok := config.Load()
	if !ok {
		log.Fatalln("環境変数から設定が取得できませんでした。終了します。")
		return
	}
	log.Println(config.SlackURL)
}
