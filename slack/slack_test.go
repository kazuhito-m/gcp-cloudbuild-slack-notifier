package slack

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/slack"
	"testing"
)

func TestSample(t *testing.T) {
	actual := slack.Sample()
	if actual != "これを返すことになっている。" {
		t.Error("結果の期待値がうまくいかなかった。", actual)
	}
}
