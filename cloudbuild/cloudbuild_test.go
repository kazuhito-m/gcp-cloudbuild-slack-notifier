package cloudbuild_test

import (
	"io/ioutil"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"os"
	"testing"
)

func Test成功時JSONの内容がParseでオブジェクト化出来る(t *testing.T) {
	jsonText := loadTestJsonText("success.json")

	actual := cloudbuild.Parse(jsonText)

	assertEquals(actual.ID, "7c91336f-90e1-4efd-aef7-53904ef0ef5f", "Jsonパーズに失敗,ID", t)
	assertEquals(actual.ProjectID, "test-development", "Jsonパーズに失敗,ProjectID", t)
	assertEquals(actual.Status, "SUCCESS", "Jsonパーズに失敗,Status", t)
}

func Test失敗時JSONの内容がParseでオブジェクト化出来る(t *testing.T) {
	jsonText := loadTestJsonText("fail.json")

	actual := cloudbuild.Parse(jsonText)

	assertEquals(actual.ID, "5f917010-9be7-438b-82d9-e90f70a8225d", "Jsonパーズに失敗,ID", t)
	assertEquals(actual.ProjectID, "test-development", "Jsonパーズに失敗,ProjectID", t)
	assertEquals(actual.Status, "FAILURE", "Jsonパーズに失敗,Status", t)
}

func TestJSONの結果をメソッドで判定出来る(t *testing.T) {
	actual := cloudbuild.Parse(loadTestJsonText("fail.json"))
	if !actual.Ng() {
		t.Error("失敗なのに成功判定になっていた。")
	}
	actual2 := cloudbuild.Parse(loadTestJsonText("success.json"))
	if !actual2.Ok() {
		t.Error("成功なのに失敗判定になっていた。")
	}
}

func assertEquals(actual string, expect string, error_message string, t *testing.T) {
	t.Helper()
	if actual != expect {
		t.Error(error_message, "actual:"+actual, "expect:"+expect)
	}
}

func loadTestJsonText(fileName string) []byte {
	dir, _ := os.Getwd()
	filePath := dir + "/testdata/" + fileName
	return loadTestFile(filePath)
}

func loadTestFile(filePath string) []byte {
	contents, _ := ioutil.ReadFile(filePath)
	return contents
}
