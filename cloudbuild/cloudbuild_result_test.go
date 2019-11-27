package cloudbuild

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := cloudbuild.Parse("{\"id\": \"abcdefg\"}")
	actual := parsed.Id
	expected := "abcdefg"

	if actual != expected {
		t.Error("JSONのパースに失敗。", actual)
	}
}
