package notify

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
)

func CreateNotify(result cloudbuild.CloudBuildResult, conf config.Config) (slack.SlackNotify, bool) {
	if result.IsStart() {
		return createStartNotify(result, conf), true
	}
	if result.IsEnd() {
		return createEndNotify(result, conf), true
	}
	return slack.SlackNotify{}, false
}

func createStartNotify(result cloudbuild.CloudBuildResult, conf config.Config) slack.SlackNotify {
	slackNotify := createBaseNotify(result, conf)

	fields := []slack.SlackField{
		slack.SlackField{
			Title: "Build ID",
			Value: "<" + result.BuildConsoleUrl() + "|" + result.ID + ">",
		},
		slack.SlackField{
			Title: "Trigger ID",
			Value: "<" + result.TriggerConsoleUrl() + "|" + result.BuildTriggerID + ">",
		},
	}

	attachement := slack.Attachment{
		Title:     makeAttachmentTitle(result),
		TitleLink: result.BuildConsoleUrl(),
		Fields:    fields,
	}

	slackNotify.Text = "CloudBuildの実行を開始しました。"
	slackNotify.Attachments = []slack.Attachment{attachement}

	return slackNotify
}
func createEndNotify(result cloudbuild.CloudBuildResult, conf config.Config) slack.SlackNotify {
	slackNotify := createBaseNotify(result, conf)

	field := slack.SlackField{
		Title: "Status",
		Value: result.Status,
	}

	attachement := slack.Attachment{
		Title:     makeAttachmentTitle(result),
		TitleLink: result.BuildConsoleUrl(),
		Color:     resultColorOf(result),
		Fields:    []slack.SlackField{field},
	}

	slackNotify.Text = makeResultText(result)
	slackNotify.Attachments = []slack.Attachment{attachement}

	return slackNotify
}

func resultColorOf(result cloudbuild.CloudBuildResult) string {
	if result.Ok() {
		return "good"
	}
	return "danger"
}

func makeResultText(result cloudbuild.CloudBuildResult) string {
	resText := "成功"
	if result.Ng() {
		resText = "失敗"
	}
	return "CloudBuildの実行が" + resText + "しました。"
}

func makeAttachmentTitle(result cloudbuild.CloudBuildResult) string {
	return result.RepositoryName() + " のビルド(" + result.BuildTriggerID + ")"
}

func createBaseNotify(result cloudbuild.CloudBuildResult, conf config.Config) slack.SlackNotify {
	slackNotify := slack.SlackNotify{}

	slackNotify.Username = "GCP CloudBuild Notification"
	slackNotify.IconURL = "https://developers.cyberagent.co.jp/blog/wp-content/uploads/2018/08/Container-Builder.png"
	slackNotify.Channel = conf.SlackChannel
	slackNotify.Mrkdwn = true

	return slackNotify
}
