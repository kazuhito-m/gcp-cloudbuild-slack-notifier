package notify

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/config"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
	"strings"
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

	fields := createBaseInfoFields(result)

	attachement := slack.Attachment{
		Title:     makeAttachmentTitle(result, conf),
		TitleLink: result.BuildConsoleUrl(),
		Fields:    fields,
	}

	slackNotify.Text = "CloudBuildの実行を *開始* しました。" + makeCustomStartMessage(result, conf)
	slackNotify.Attachments = []slack.Attachment{attachement}

	return slackNotify
}
func createEndNotify(result cloudbuild.CloudBuildResult, conf config.Config) slack.SlackNotify {
	slackNotify := createBaseNotify(result, conf)

	fields := []slack.SlackField{
		fieldOf("Status", "*"+result.Status+"*"),
		fieldOf("Total Time", result.TotalTime()),
	}
	fields = append(fields, createBaseInfoFields(result)...)
	if len(result.ErrorSteps()) > 0 {
		fields = append(fields, createErrorStepField(result.ErrorSteps()[0]))
	}

	attachement := slack.Attachment{
		Title:     makeAttachmentTitle(result, conf),
		TitleLink: result.BuildConsoleUrl(),
		Color:     resultColorOf(result),
		Fields:    fields,
	}

	slackNotify.Text = makeResultText(result, conf)
	slackNotify.Attachments = []slack.Attachment{attachement}

	return slackNotify
}

func resultColorOf(result cloudbuild.CloudBuildResult) string {
	if result.Ok() {
		return "good"
	}
	return "danger"
}

func makeResultText(result cloudbuild.CloudBuildResult, conf config.Config) string {
	resText := "成功"
	if result.Ng() {
		resText = "失敗"
	}
	return "CloudBuildの実行が *" + resText + "* しました。" + makeCustomEndMessage(result, conf)
}

func makeCustomStartMessage(result cloudbuild.CloudBuildResult, conf config.Config) string {
	customMessage := ""
	triggerSetting, hit := conf.TriggerSettingOf(result.BuildTriggerID)
	if hit {
		customMessage = triggerSetting.CustumStartMessage
	}
	if customMessage == "" {
		return ""
	}
	return "\n\n" + customMessage
}

func makeCustomEndMessage(result cloudbuild.CloudBuildResult, conf config.Config) string {
	customMessage := ""
	triggerSetting, hit := conf.TriggerSettingOf(result.BuildTriggerID)
	if hit {
		customMessage = triggerSetting.CustumEndMessage
	}
	if customMessage == "" {
		return ""
	}
	return "\n\n" + customMessage
}

func makeAttachmentTitle(result cloudbuild.CloudBuildResult, conf config.Config) string {
	triggerSetting, hit := conf.TriggerSettingOf(result.BuildTriggerID)
	if hit {
		return triggerSetting.AliasName
	}
	return result.RepositoryName() + " のビルド"
}

func createBaseNotify(result cloudbuild.CloudBuildResult, conf config.Config) slack.SlackNotify {
	slackNotify := slack.SlackNotify{}

	slackNotify.Username = "GCP CloudBuild Notification"
	slackNotify.IconURL = "https://developers.cyberagent.co.jp/blog/wp-content/uploads/2018/08/Container-Builder.png"
	slackNotify.Channel = makeChannelName(result, conf)
	slackNotify.Mrkdwn = true

	return slackNotify
}

func makeChannelName(result cloudbuild.CloudBuildResult, conf config.Config) string {
	triggerSetting, hit := conf.TriggerSettingOf(result.BuildTriggerID)
	customSlackChannel := ""
	if hit {
		customSlackChannel = triggerSetting.SlackChannel
	}
	if customSlackChannel != "" {
		return customSlackChannel
	}
	return conf.SlackChannel
}

func createBaseInfoFields(result cloudbuild.CloudBuildResult) []slack.SlackField {
	return []slack.SlackField{
		fieldOf("Project ID", result.ProjectID),
		fieldOf("Source Repository/Branch(or Tag)Name", makeSourceRepositoryLink(result)),
		fieldOf("Build ID", linkOf(result.ID, result.BuildConsoleUrl())),
		fieldOf("Trigger ID", linkOf(result.BuildTriggerID, result.TriggerConsoleUrl())),
	}
}

func makeSourceRepositoryLink(result cloudbuild.CloudBuildResult) string {
	tagOrBranchName := result.BranchName()
	tagName := result.TagName()
	if tagName != "" {
		tagOrBranchName = tagName
	}

	nameAndBranch := result.RepositoryName() + "/" + tagOrBranchName
	if tagName != "" {
		nameAndBranch += "(Tag)"
	}

	if result.UseOutsideSouceRepository() {
		return nameAndBranch
	}
	src := result.Source.RepoSource
	url := "https://source.cloud.google.com/" + src.ProjectID + "/" + src.RepoName + "/+/" + tagOrBranchName + ":"
	return linkOf(nameAndBranch, url)
}
func linkOf(caption string, url string) string {
	return "<" + url + "|" + caption + ">"
}

func fieldOf(title string, value string) slack.SlackField {
	return slack.SlackField{
		Title: title,
		Value: value,
		Short: true,
	}
}

func createErrorStepField(errorStep cloudbuild.CloudBuildStep) slack.SlackField {
	content := " * ID/Name: " + errorStep.Description()
	content += "\n * Status: " + errorStep.Status
	content += "\n * EntryPoint: `" + errorStep.Entrypoint + "`"
	content += "\n * Args: ```" + strings.Join(errorStep.Args, "\n") + "```"
	return slack.SlackField{
		Title: "Error Step",
		Value: content,
		Short: false,
	}
}
