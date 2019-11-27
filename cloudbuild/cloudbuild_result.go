package cloudbuild

import (
	"encoding/json"
	"log"
)

type CluodBuildResult struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	Source    Source `json:"source"`
	Steps     []Step `json:"steps"`
}

type Source struct {
	RepoSource RepoSource `json:"repoSource"`
}

type RepoSource struct {
	ProjectId  string `json:"projectId"`
	RepoName   string `json:"repoName"`
	BranchName string `json:"branchName"`
}

type Step struct {
	Name string `json:"name`
	// "args":" [
	// 	"-c",
	// 	"echo 'この日本語が出力されたら成功とする'\n"
	// ],
	// "id":" "test-steps-01",
	// "entrypoint":" "/bin/bash",
	// "timing":" {
	// 	"startTime":" "2019-11-21T06:"02:"47.410751454Z",
	// 	"endTime":" "2019-11-21T06:"02:"49.444260195Z"
	// },
	// "pullTiming":" {
	// 	"startTime":" "2019-11-21T06:"02:"47.410751454Z",
	// 	"endTime":" "2019-11-21T06:"02:"48.592274335Z"
	// },
	// "status":" "SUCCESS"
}

func Parse(json_text string) CluodBuildResult {
	bytes := []byte(json_text)
	var decode_data CluodBuildResult
	if err := json.Unmarshal(bytes, &decode_data); err != nil {
		log.Fatal(err)
	}
	return decode_data
}
