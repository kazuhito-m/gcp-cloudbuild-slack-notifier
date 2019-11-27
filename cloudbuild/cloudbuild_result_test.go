package cloudbuild

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := cloudbuild.Parse("{\"id\": \"abcdefg\"}")
	assert(parsed.Id, "abcdefg", "JSONのパーズに失敗。", t)
}

func assert(actual string, expect string, error_message string, t *testing.T) {
	if actual != expect {
		t.Error(error_message, "actual:", actual, ",expect:", expect)
	}
}
