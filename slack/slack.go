package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"log"
	"net/http"
)

func Sample() string {
	fmt.Printf("testtest\n")
	return "これを返すことになっている。"
}

func SendNotify(notify SlackNotify, configure config.Config) bool {
	jsonBytes, _ := json.Marshal(notify)
	jsonStr := string(jsonBytes)

	log.Println("Send contents(json):" + jsonStr)

	req, err := http.NewRequest("POST", configure.SlackURL, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Println(err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println(resp)
	defer resp.Body.Close()

	return true
}