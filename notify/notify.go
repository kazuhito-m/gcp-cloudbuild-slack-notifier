package notify

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
)

func CreateNotify(cloudBuildResult cloudbuild.CloudBuildResult) slack.SlackNotify {
	slackNotify := slack.SlackNotify{}

	field := slack.SlackField{
		Title: "Status",
		Value: cloudBuildResult.Status,
	}

	attachement := slack.Attachment{
		Title:     "タイトル",
		TitleLink: "https://google.com",
		Fields:    []slack.SlackField{field},
	}

	slackNotify.Username = "GCP CloudBuild Notification"
	slackNotify.IconURL = "https://developers.cyberagent.co.jp/blog/wp-content/uploads/2018/08/Container-Builder.png"
	slackNotify.Channel = "#hanguottest"
	slackNotify.Text = "テスト"
	slackNotify.Mrkdwn = true
	slackNotify.Attachments = []slack.Attachment{attachement}

	return slackNotify
}
