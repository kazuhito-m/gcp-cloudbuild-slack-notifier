package cloudbuild_test

import (
	"encoding/json"
	"kazuhito-m/gcp-cloudbuild-slack-notifier/cloudbuild"
	"os"
	"testing"
)

func TestソースリポジトリがGCP内かそれ以外かを判定できる(t *testing.T) {
	outsideResult := loadTestCloudBuildResult("gcp_outside.json")
	actual := outsideResult.UseOutsideSouceRepository()
	assertTrue(actual, "GCP外のSourceRepositoryなのに、判定できなかった。", t)

	outsideResult2 := loadTestCloudBuildResult("gcp_inside.json")
	actual2 := outsideResult2.UseOutsideSouceRepository()
	assertFalse(actual2, "GCP内のSourceRepositoryなのに、判定できなかった。", t)
}

func assertTrue(actual bool, errorMessage string, t *testing.T) {
	t.Helper()
	if !actual {
		t.Error(errorMessage)
	}
}

func assertFalse(actual bool, errorMessage string, t *testing.T) {
	t.Helper()
	if actual {
		t.Error(errorMessage)
	}
}

func loadTestCloudBuildResult(fileName string) cloudbuild.CloudBuildResult {
	dir, _ := os.Getwd()
	filePath := dir + "/testdata/" + fileName
	bytes := loadTestFile(filePath)
	var decode_data cloudbuild.CloudBuildResult
	json.Unmarshal(bytes, &decode_data)
	return decode_data
}
