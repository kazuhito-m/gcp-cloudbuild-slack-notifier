package cloudbuild

import (
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"testing"
)

func TestParse(t *testing.T) {
	parsed := cloudbuild.Parse("{\"id\": \"abcdefg\"}")
	assertEquals(parsed.Id, "abcdefg", "JSONのパーズに失敗。", t)
}

func assertEquals(actual string, expect string, error_message string, t *testing.T) {
	t.Helper()
	if actual != expect {
		t.Error(error_message, "actual:", actual, ",expect:", expect)
	}
}
