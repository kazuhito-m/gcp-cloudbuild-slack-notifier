package pubsub

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/pubsub"
	"testing"
)

func TestJsonからstatus文字列だけを取り出す(t *testing.T) {
	jsonText := `{"id": "test", "status": "FAILURE", "steps": [{ "status": "SUCCESS"}]}`

	actual := pubsub.PickUpStatusText(jsonText)

	assertEquals(actual, "FAILURE", "json内からstatus文字列を取得できなかった。", t)
}

func assertEquals(actual string, expect string, error_message string, t *testing.T) {
	t.Helper()
	if actual != expect {
		t.Error(error_message, "actual:"+actual, "expect:"+expect)
	}
}
